vim.pack.add({
    { src = "https://github.com/dgox16/oldworld.nvim" }
})

require("oldworld").setup({
    variant = "oled",
})

vim.cmd("colorscheme oldworld")
