# omz
ZSH_THEME="sorin"
plugins=(git)
if [ -f $ZDOTDIR/ohmyzsh/oh-my-zsh.sh ]; then
    source $ZDOTDIR/ohmyzsh/oh-my-zsh.sh
elif [ -f ~/.oh-my-zsh/oh-my-zsh.sh ]; then
    source ~/.oh-my-zsh/oh-my-zsh.sh
fi

# aliases
alias modify-repos="sudo reflector --save /etc/pacman.d/mirrorlist --protocol https --latest 5 --sort age"

alias rcs-cloud="rclone -v sync ~/Cloud mega-crypt:"
alias rcs-local="rclone -v sync mega-crypt: ~/Cloud"
alias rc-check="rclone -v check mega-crypt: ~/Cloud"

alias todo="$EDITOR $TODO_FILE"
alias rmbadlinks="find . -xtype l -delete"

alias ws="windscribe-cli"
alias pv="protonvpn"

# vi mode
bindkey -v

export KEYTIMEOUT=1
setopt PROMPT_SUBST

function vi_mode_status() {
  if [[ ${KEYMAP} == vicmd ]]; then
    echo "%F{cyan}[NORMAL]%f"
  else
    echo "%F{green}[INSERT]%f"
  fi
}

function zle-keymap-select {
  if [[ ${KEYMAP} == vicmd ]]; then
    echo -ne '\e[1 q'
  else
    echo -ne '\e[5 q'
  fi
  zle reset-prompt
}
zle -N zle-keymap-select

function zle-line-init {
  echo -ne '\e[5 q'
  zle reset-prompt
}
zle -N zle-line-init

RPS1="$RPS1 "'$(vi_mode_status)'

# fzf
source <(fzf --zsh)
export FZF_DEFAULT_OPTS="--color=base16"
export FZF_DEFAULT_COMMAND="fd --type f"
export FZF_CTRL_T_COMMAND="fd --type f"
export FZF_ALT_C_COMMAND="fd --type d"

# tmux new session
# if [ -x "$(command -v tmux)" ] && [ -n "${DISPLAY}" ] && [ -z "${TMUX}" ]; then
#     exec tmux new-session >/dev/null 2>&1
# fi

# do not allow ctrl d to kill session
setopt IGNOREEOF
