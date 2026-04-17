vim.pack.add({
    { src = "https://github.com/nvim-treesitter/nvim-treesitter" },
})

require("nvim-treesitter").setup({
})

vim.api.nvim_create_autocmd('FileType', {
    callback = function()
        pcall(vim.treesitter.start)
        vim.bo.indentexpr = "v:lua.require'nvim-treesitter'.indentexpr()"
    end,
})

local ensureInstalled = {
    "bash",
    "c",
    "cpp",
    "css",
    "git_config",
    "git_rebase",
    "gitcommit",
    "gitignore",
    "go",
    "html",
    "htmldjango",
    "hyprlang",
    "javascript",
    "jinja",
    "jinja_inline",
    "json",
    "lua",
    "markdown",
    "markdown_inline",
    "php",
    "python",
    "rust",
    "scss",
    "sql",
    "tmux",
    "vim",
    "vimdoc",
    "xml",
}

local alreadyInstalled = require('nvim-treesitter.config').get_installed()
local parsersToInstall = vim.iter(ensureInstalled)
    :filter(function(parser)
        return not vim.tbl_contains(alreadyInstalled, parser)
    end)
    :totable()
require('nvim-treesitter').install(parsersToInstall)
