#/bin/bash

function console_blue() {
    echo -e "\033[36m[ $1 ]\033[0m"
}

function console_green() {
    echo -e "\033[32m[ $1 ]\033[0m"
}

function console_orangered() {
    echo -e "\033[31m\033[01m[ $1 ]\033[0m"
}

function console_yellow() {
    echo -e "\033[33m\033[01m[ $1 ]\033[0m"
}

function auto_generate_menu() {
  git pull origin main
  ./script/spider-blog-title
  rm -rf source/menu/index.md
  mv ./blog-menu.md source/menu/index.md
  nvm use 16.2.0
  hexo cl
  hexo g
  hexo d
  git status
  git add -A
  git commit -m "auto generate menu"-`date +"%Y-%m-%d_%H:%M:%S"`
  git push -u origin main
}

auto_generate_menu
