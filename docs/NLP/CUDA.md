# CUDA

[VSC插件](https://marketplace.visualstudio.com/items?itemName=NVIDIA.nsight-vscode-edition)

## 在 Windows 系统安装

如果使用 Windows 系统，CUDA 依赖于 [Visual Studio](https://visualstudio.microsoft.com/) 的 C++ 编译器。

然后在该 [链接](https://developer.nvidia.com/cuda-downloads) 下载并安装 CUDA。

安装完成后执行 `nvcc --version` 和 `nvidia-smi` 检查。

## 在 Arch Linux 安装

如果使用 Arch Linux, 要注意使用 Nvidia 的官方驱动，不要使用开源的 nouveau 驱动。

```bash
sudo pacman -Syu nvidia nvidia-utils nvidia-settings
sudo reboot
nvidia-smi

sudo pacman -S cuda
```

配置环境变量

```bash
export PATH=/opt/cuda/bin:$PATH
export LD_LIBRARY_PATH=/opt/cuda/lib64:$LD_LIBRARY_PATH
source ~/.bashrc   # 或 ~/.zshrc
```

可选工具

| 工具             | 功能                                  |
| -------------- | ----------------------------------- |
| `cuda-samples` | 安装后可用 `make` 编译示例程序                 |
| `nvtop`        | 实时查看 GPU 利用率：`sudo pacman -S nvtop` |
| `nsight`       | NVIDIA 开发工具链，图形化调试与性能分析             |
| `yay -S cudnn` | 安装 CUDNN 加速库（如有深度学习需求）              |

## 编程指南

https://developer.nvidia.com/blog/easy-introduction-cuda-c-and-c/

https://docs.nvidia.com/cuda/cuda-c-programming-guide/
