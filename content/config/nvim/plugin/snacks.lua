vim.pack.add({
    { src = "https://github.com/folke/snacks.nvim" },
    { src = "https://github.com/nvim-tree/nvim-web-devicons" }
})

require("snacks").setup({
    explorer = { enabled = true },
    picker = {
        enabled = true,
        layout = {
            layout = { backdrop = false }
        },
        sources = {
            files = { hidden = true },
            grep = { hidden = true },
            explorer = { hidden = true }
        },
    },
    scroll = { enabled = true },
    notifier = { enabled = true }
})
