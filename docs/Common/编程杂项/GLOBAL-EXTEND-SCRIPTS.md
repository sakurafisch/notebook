# GLOBAL EXTEND SCRIPTS

>  适用于 v2.5.10

## 保持简单

```javascript
// ============================================================
// 基于 Loyalsoldier 规则集：国内直连、国外代理
// ============================================================

function main(config, profileName) {
  // ---------- 1. 定义 Loyalsoldier 规则集（rule-providers） ----------
  const ruleProviders = {
    reject: {
      type: "http",
      behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/reject.txt",
      path: "./ruleset/reject.yaml",
      interval: 86400
    },
    icloud: {
      type: "http",
      behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/icloud.txt",
      path: "./ruleset/icloud.yaml",
      interval: 86400
    },
    apple: {
      type: "http",
      behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/apple.txt",
      path: "./ruleset/apple.yaml",
      interval: 86400
    },
    direct: {
      type: "http",
      behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/direct.txt",
      path: "./ruleset/direct.yaml",
      interval: 86400
    },
    proxy: {
      type: "http",
      behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/proxy.txt",
      path: "./ruleset/proxy.yaml",
      interval: 86400
    },
    private: {
      type: "http",
      behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/private.txt",
      path: "./ruleset/private.yaml",
      interval: 86400
    },
    gfw: {
      type: "http",
      behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/gfw.txt",
      path: "./ruleset/gfw.yaml",
      interval: 86400
    },
    "tld-not-cn": {
      type: "http",
      behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/tld-not-cn.txt",
      path: "./ruleset/tld-not-cn.yaml",
      interval: 86400
    },
    telegramcidr: {
      type: "http",
      behavior: "ipcidr",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/telegramcidr.txt",
      path: "./ruleset/telegramcidr.yaml",
      interval: 86400
    },
    cncidr: {
      type: "http",
      behavior: "ipcidr",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/cncidr.txt",
      path: "./ruleset/cncidr.yaml",
      interval: 86400
    },
    lancidr: {
      type: "http",
      behavior: "ipcidr",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/lancidr.txt",
      path: "./ruleset/lancidr.yaml",
      interval: 86400
    }
  };

  // 合并到 config 中（保留订阅自带的 rule-providers）
  config["rule-providers"] = {
    ...(config["rule-providers"] || {}),
    ...ruleProviders
  };

  // ---------- 2. 动态探测代理组名称（避免硬编码导致切换订阅报错） ----------
  // 参考来源：https://locdd.com/t/topic/104662 的自适应策略组思路
  const groups = (config["proxy-groups"] || []).map(g => g.name);
  const targetGroup = groups.find(name =>
    name.includes("节点选择") ||
    name.includes("PROXY") ||
    name.includes("Proxy") ||
    name.includes("手动")
  ) || groups[0] || "DIRECT"; // 兜底：如果找不到任何组，就直连（避免配置校验失败）

  // ---------- 3. 定义 Loyalsoldier 规则列表（rules） ----------
  // 顺序严格遵循 Loyalsoldier README 的推荐顺序
  const loyalsoldierRules = [
    "RULE-SET,private,DIRECT",
    "RULE-SET,reject,REJECT",
    "RULE-SET,icloud,DIRECT",
    "RULE-SET,apple,DIRECT",
    "RULE-SET,direct,DIRECT",
    "RULE-SET,lancidr,DIRECT",
    "RULE-SET,cncidr,DIRECT",
    `RULE-SET,telegramcidr,${targetGroup}`,
    `RULE-SET,proxy,${targetGroup}`,
    `RULE-SET,gfw,${targetGroup}`,
    `RULE-SET,tld-not-cn,${targetGroup}`,
    `MATCH,${targetGroup}`
  ];

  // ---------- 4. 注入到规则最前端（前置），并防止重复注入 ----------
  if (!config.rules) {
    config.rules = [];
  }

  // 逆序 unshift，确保最终顺序与 loyalsoldierRules 数组一致
  for (let i = loyalsoldierRules.length - 1; i >= 0; i--) {
    const rule = loyalsoldierRules[i];
    if (!config.rules.includes(rule)) {
      config.rules.unshift(rule);
    }
  }

  return config;
}

```

## 详细定义

```javascript
// ============================================================
// 基于 Loyalsoldier 规则集
// 切换黑白名单：只改 MODE 变量
// ============================================================

// 危险内置组：不能作为代理出口
const DANGEROUS_NAMES = new Set(["REJECT", "REJECT-DROP"]);

// 脚本管理的规则集名单，用于注入前清理旧规则
const OUR_RULESETS = new Set([
  "private", "reject", "icloud", "apple", "direct",
  "lancidr", "cncidr", "telegramcidr", "proxy", "gfw", "tld-not-cn"
]);

// console 安全包装
function safeLog(level, msg) {
  try {
    if (typeof console !== "undefined" && typeof console[level] === "function") {
      console[level]("[Loyalsoldier] " + msg);
    }
  } catch (e) {
    // 忽略日志异常
  }
}

// 两位补零
function pad2(n) {
  return n < 10 ? "0" + n : String(n);
}

function main(config, profileName) {
  // ============================================================
  // ★★★ 唯一需要修改的变量 ★★★
  // "whitelist" = 白名单：国内直连，其余全部走代理
  // "blacklist" = 黑名单：默认直连，仅 gfw/proxy 等名单走代理
  const MODE = "whitelist";
  // ============================================================

  // ---------- 1. 定义并合并 Loyalsoldier 规则集 ----------
  const ruleProviders = {
    reject: {
      type: "http", behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/reject.txt",
      path: "./ruleset/reject.yaml", interval: 86400
    },
    icloud: {
      type: "http", behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/icloud.txt",
      path: "./ruleset/icloud.yaml", interval: 86400
    },
    apple: {
      type: "http", behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/apple.txt",
      path: "./ruleset/apple.yaml", interval: 86400
    },
    direct: {
      type: "http", behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/direct.txt",
      path: "./ruleset/direct.yaml", interval: 86400
    },
    proxy: {
      type: "http", behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/proxy.txt",
      path: "./ruleset/proxy.yaml", interval: 86400
    },
    private: {
      type: "http", behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/private.txt",
      path: "./ruleset/private.yaml", interval: 86400
    },
    gfw: {
      type: "http", behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/gfw.txt",
      path: "./ruleset/gfw.yaml", interval: 86400
    },
    "tld-not-cn": {
      type: "http", behavior: "domain",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/tld-not-cn.txt",
      path: "./ruleset/tld-not-cn.yaml", interval: 86400
    },
    telegramcidr: {
      type: "http", behavior: "ipcidr",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/telegramcidr.txt",
      path: "./ruleset/telegramcidr.yaml", interval: 86400
    },
    cncidr: {
      type: "http", behavior: "ipcidr",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/cncidr.txt",
      path: "./ruleset/cncidr.yaml", interval: 86400
    },
    lancidr: {
      type: "http", behavior: "ipcidr",
      url: "https://cdn.jsdelivr.net/gh/Loyalsoldier/clash-rules@release/lancidr.txt",
      path: "./ruleset/lancidr.yaml", interval: 86400
    }
  };

  config["rule-providers"] = Object.assign(
    {},
    config["rule-providers"] || {},
    ruleProviders
  );

  // ---------- 2. 解析代理出口 ----------
  const targetGroup = resolveTargetGroup(config);

  // ---------- 3. 按 MODE 构建规则列表 ----------
  const directRules = [
    "RULE-SET,private,DIRECT",
    "RULE-SET,reject,REJECT",
    "RULE-SET,icloud,DIRECT",
    "RULE-SET,apple,DIRECT",
    "RULE-SET,direct,DIRECT",
    "RULE-SET,lancidr,DIRECT",
    "RULE-SET,cncidr,DIRECT"
  ];

  const proxyRules = [
    "RULE-SET,telegramcidr," + targetGroup,
    "RULE-SET,proxy," + targetGroup,
    "RULE-SET,gfw," + targetGroup,
    "RULE-SET,tld-not-cn," + targetGroup
  ];

  const tailRule = MODE === "whitelist"
    ? "MATCH," + targetGroup
    : "MATCH,DIRECT";

  const loyalsoldierRules = directRules.concat(proxyRules, [tailRule]);

  // ---------- 4. 清理上一轮注入的旧规则 ----------
  if (!Array.isArray(config.rules)) {
    config.rules = [];
  }

  config.rules = config.rules.filter(function (rule) {
    if (typeof rule !== "string") return true;
    if (rule === "MATCH" || rule.indexOf("MATCH,") === 0) return false;
    if (rule.indexOf("RULE-SET,") !== 0) return true;
    const parts = rule.split(",");
    const setName = parts[1] ? parts[1].trim() : "";
    return !OUR_RULESETS.has(setName);
  });

  // ---------- 5. 前置注入 ----------
  for (let i = loyalsoldierRules.length - 1; i >= 0; i--) {
    config.rules.unshift(loyalsoldierRules[i]);
  }

  return config;
}

// ============================================================
// 解析代理出口
//
// 【新】选择逻辑：
//   1. groups 非空时，默认选择第一个（groups[0]）
//   2. 如果第一个是 "直连" 或 "DIRECT"（不区分大小写）：
//      - 存在第二个组 → 选择下一个
//      - 不存在第二个组 → 放弃重选，保持第一个
//   3. 安全检查：跳过 REJECT / REJECT-DROP 等危险组，继续往后找
//   4. groups 为空时，收集 proxies / proxy-providers 动态创建
//      AUTO-PROXY-YYYYMMDD-HHMMSS（带日期时间戳，避免重名）
// ============================================================
function resolveTargetGroup(config) {
  const groups = Array.isArray(config["proxy-groups"]) ? config["proxy-groups"] : [];

  if (groups.length > 0) {
    let selectedIndex = 0;
    const firstName = (groups[0].name || "").trim();
    const isDirect = firstName.toUpperCase() === "DIRECT" || firstName === "直连";

    // 【新】默认选择第一个；如果第一个是 DIRECT/直连，尝试选择下一个
    if (isDirect && groups.length > 1) {
      selectedIndex = 1;
      safeLog("log", "第一个代理组为 " + firstName + "，已跳过并选择下一个");
    } else if (isDirect) {
      safeLog("warn", "唯一代理组为 " + firstName + "，放弃重选");
    }

    // 安全检查：跳过危险组（REJECT / REJECT-DROP）
    while (selectedIndex < groups.length) {
      const name = (groups[selectedIndex].name || "").trim().toUpperCase();
      if (!DANGEROUS_NAMES.has(name)) break;
      selectedIndex++;
    }

    if (selectedIndex < groups.length) {
      return groups[selectedIndex].name;
    }

    // 所有候选组都是危险组，转走 AUTO-PROXY 创建流程
    safeLog("warn", "候选组均为危险组，尝试创建 AUTO-PROXY");
  }

  // groups 为空（或全是危险组）：收集节点来源，创建 AUTO-PROXY
  const rawProxies = Array.isArray(config.proxies) ? config.proxies : [];
  const proxyNames = rawProxies
    .map(function (p) { return typeof p === "string" ? p : (p && p.name); })
    .filter(Boolean);
  const providerKeys = Object.keys(config["proxy-providers"] || {});

  if (proxyNames.length === 0 && providerKeys.length === 0) {
    safeLog("warn", "未探测到任何代理组或节点来源，规则回退到 DIRECT");
    return "DIRECT";
  }

  // 生成带日期时间戳的名称，避免与订阅自带组重名
  const now = new Date();
  const timestamp =
    now.getFullYear() +
    pad2(now.getMonth() + 1) +
    pad2(now.getDate()) + "-" +
    pad2(now.getHours()) +
    pad2(now.getMinutes()) +
    pad2(now.getSeconds());
  const fallbackName = "AUTO-PROXY-" + timestamp;

  const fallbackGroup = { name: fallbackName, type: "select" };
  if (proxyNames.length > 0) fallbackGroup.proxies = proxyNames;
  if (providerKeys.length > 0) fallbackGroup.use = providerKeys;

  config["proxy-groups"] = [fallbackGroup].concat(groups);

  safeLog("log", "已创建代理组 " + fallbackName +
    " (proxies=" + proxyNames.length + ", providers=" + providerKeys.length + ")");
  return fallbackName;
}

```

