git status
git add -A
git commit -m "auto commit $(date)"
export http_proxy=localhost:8889
export https_proxy=localhost:8889
git push -u origin master