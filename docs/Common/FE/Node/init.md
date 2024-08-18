# 新建后端项目



```bash
npm init -y
yarn add -D @types/node typescript
yarn add -D ts-node
yarn add -D nodemon
npx tsconfig.json
```

```json
"script": {
    "watch": "tsc -w",
    "dev": "nodemon --exec ts-node src/index.ts",
    "start": "ts-node src/index.ts"
}
```

