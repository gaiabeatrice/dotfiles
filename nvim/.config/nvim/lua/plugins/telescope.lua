return {
  "nvim-telescope/telescope.nvim",
  opts = {
    pickers = {
      find_files = {
        find_command = { "rg", "--files", "--hidden", "-g", "!.git" },
      },
    },
  },
}
