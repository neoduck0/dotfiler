vim.pack.add({
    { src = "https://github.com/mason-org/mason.nvim" }
})

require("mason").setup()

local packages = {
    bash = { "bash-language-server", "shfmt" },
    c = { "clangd" },
    css = { "css-lsp" },
    go = { "gopls", "gofumpt" },
    html = { "html-lsp" },
    json = { "json-lsp", "jq" },
    lua = { "lua-language-server" },
    python = { "pyright", "ruff" },
    typescript = { "typescript-language-server" },
    php = { "intelephense" }
}

local to_install = {}

for _, pkgs in pairs(packages) do
    for _, pkg in ipairs(pkgs) do
        if not require("mason-registry").is_installed(pkg) then
            table.insert(to_install, pkg)
        end
    end
end

if #to_install > 0 then
    local registry = require("mason-registry")
    for _, pkg in ipairs(to_install) do
        local ok, err = pcall(registry.get_package(pkg).install, registry.get_package(pkg))
        if not ok then
            vim.notify("Failed to install " .. pkg .. ": " .. err, vim.log.LEVEL_ERROR)
        end
    end
end
