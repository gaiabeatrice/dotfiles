return {
  {
    "vim-dadbod-ui",
    init = function()
      vim.g.db_ui_use_nerd_fonts = true
      vim.g.db_ui_execute_on_save = false
      vim.g.db_ui_disable_mappings = true
      vim.g.db_ui_save_location = "./gaias_queries"
      vim.g.db_ui_use_nvim_notify = true
    end,
    lazy = false,
  },
}
