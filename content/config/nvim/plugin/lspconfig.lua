vim.pack.add({
    { src = "https://github.com/neovim/nvim-lspconfig" },
    { src = "https://github.com/folke/lazydev.nvim" },
})

require("lazydev").setup({
    library = {
        path = "${3rd}/luv/library",
        words = { "vim%.uv" }
    }
})

vim.lsp.config("pyright", {
    settings = {
        python = {
            pythonPath = ".venv/bin/python"
        }
    }
})

vim.lsp.enable("bashls")
vim.lsp.enable("clangd")
vim.lsp.enable("cssls")
vim.lsp.enable("gopls")
vim.lsp.enable("html")
vim.lsp.enable("lua_ls")
vim.lsp.enable("pyright")
vim.lsp.enable("ts_ls")
vim.lsp.enable("jsonls")
vim.lsp.enable("intelephense")

vim.diagnostic.config({
    virtual_text = {
        current_line = true,
    },
    signs = false,
    severity_sort = true
})
