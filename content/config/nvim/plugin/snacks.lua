vim.pack.add({
    { src = "https://github.com/folke/snacks.nvim" },
    { src = "https://github.com/nvim-tree/nvim-web-devicons" }
})

require("snacks").setup({
    dashboard = {
        enabled = true,
        sections = {
            { section = "header" },
            { section = "keys",  gap = 1, padding = 1 },
            -- { section = "recent_files", icon = " ", title = "Recent Files", padding = 1 },
            -- { section = "startup" },
        },
        preset = {
            keys = {
                { icon = "", key = "f", desc = "Find File", action = ":lua Snacks.dashboard.pick('files')" },
                { icon = "", key = "n", desc = "New File", action = ":ene | startinsert" },
                { icon = "", key = "/", desc = "Find Text", action = ":lua Snacks.dashboard.pick('live_grep')" },
                { icon = "", key = "r", desc = "Recent Files", action = ":lua Snacks.dashboard.pick('oldfiles')" },
                { icon = "", key = "c", desc = "Config", action = ":lua Snacks.dashboard.pick('files', {cwd = vim.fn.stdpath('config')})" },
                { icon = "", key = "h", desc = "Help", action = ":lua Snacks.picker.help()" },
                { icon = "", key = "q", desc = "Quit", action = ":qa" },
            }
        }
    },
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
    notifier = { enabled = true },
    indent = {
        enabled = true,
        only_scope = false,
    },
})
