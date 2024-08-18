# Visualizations

[参考链接](https://chromium.googlesource.com/chromium/src/+/main/docs/process_model_and_site_isolation.md#visualizations)

Chromium provides several ways to view the current state of the process model:

- Chromium's Task Manager: This can be found under “More Tools” in the menu, and shows live resource usage for each of Chromium's processes. The Task Manager also shows which documents and workers are grouped together in a given process: only the first row of a given group displays process ID and most statistics, and all rows of a group are highlighted when one is clicked. Note that double clicking any row attempts to switch to the tab it is associated with. In the default sort order (i.e., when clicking the Task column header until the up/down triangle disappears), processes for subframes are listed under the process for their tab when possible, although this may not be possible if subframes from multiple tabs are in a given process.
- `chrome://process-internals/#web-contents`: This is an internal diagnostic page which shows information about the SiteInstances and processes for each open document.
- `chrome://discards/graph`: This is an internal diagnostic page that includes a visualization of how the open documents and workers map to processes. Clicking on any node provides more details.