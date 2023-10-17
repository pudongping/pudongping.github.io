#!/bin/bash
# deploy gh-pages branch

base_path=$(cd `dirname $0`; pwd)

cd $base_path

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

function deploy_2_gh_pages() {
  nvm use 16.2.0
  hexo cl
  hexo g
  hexo d
}

deploy_2_gh_pages
