vim.keymap.set("n", "<localleader>r", ":normal vip<CR><PLUG>(DBUI_ExecuteQuery)", { buffer = true })
vim.keymap.set("n", "<localleader>w", "<PLUG>(DBUI_SaveQuery)", { buffer = true })
