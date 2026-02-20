vim.pack.add({
    { src = "https://github.com/stevearc/conform.nvim" }
})

require("conform").setup({
    formatters_by_ft = {
        ["*"] = { "trim_whitespace", "trim_newlines" },
        go = { "gofumpt" },
        json = { "jq" },
        lua = { lsp_format = "first" },
        python = { "ruff" },
        sh = { "shfmt" },
        html = { lsp_format = "first" },
    },
    format_on_save = {
        timeout_ms = 500,
    },
})
