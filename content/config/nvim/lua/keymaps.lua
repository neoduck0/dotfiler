vim.g.mapleader = " "

vim.keymap.set("n", "<leader>e", function() Snacks.explorer() end)

vim.keymap.set("n", "<leader>tw", function()
    vim.opt.wrap = not vim.opt.wrap:get()
    vim.opt.linebreak = vim.opt.wrap:get()
end)
vim.keymap.set('n', '<leader>td', function()
    vim.diagnostic.enable(not vim.diagnostic.is_enabled())
end)
vim.keymap.set('n', '<leader>ts', function()
    vim.opt.spell = not vim.opt.spell:get()
end)


vim.keymap.set("n", "<leader>f/", function() Snacks.picker.grep() end)
vim.keymap.set("n", "<leader>fD", function() Snacks.picker.diagnostics() end)
vim.keymap.set("n", "<leader>fb", function() Snacks.picker.buffers() end)
vim.keymap.set("n", "<leader>fd", function() Snacks.picker.diagnostics_buffer() end)
vim.keymap.set("n", "<leader>ff", function() Snacks.picker.files() end)
vim.keymap.set("n", "<leader>fh", function() Snacks.picker.help() end)
vim.keymap.set("n", "<leader>fp", function() Snacks.picker.projects() end)
vim.keymap.set("n", "<leader>fg", function() Snacks.picker.git_diff() end)
vim.keymap.set("n", "<leader>fG", function() Snacks.picker.git_status() end)


vim.keymap.set('n', '[c', function() require("gitsigns").nav_hunk('prev') end)
vim.keymap.set('n', ']c', function() require("gitsigns").nav_hunk('next') end)
