return {
  {
    "christoomey/vim-tmux-navigator",
    cmd = {
      "TmuxNavigateLeft",
      "TmuxNavigateDown",
      "TmuxNavigateUp",
      "TmuxNavigateRight",
      "TmuxNavigatePrevious",
    },
    keys = {
      { "<c-h>", "<cmd><C-U>TmuxNavigateLeft<cr>" },
      { "<c-j>", "<cmd><C-U>TmuxNavigateDown<cr>" },
      { "<c-k>", "<cmd><C-U>TmuxNavigateUp<cr>" },
      { "<c-l>", "<cmd><C-U>TmuxNavigateRight<cr>" },
      { "<c-\\>", "<cmd><C-U>TmuxNavigatePrevious<cr>" },
    },
  },
  {
    "janko/vim-test",
    dependencies = {
      "tpope/vim-dispatch",
      "preservim/vimux",
    },
    event = "VimEnter",
    config = function()
      vim.g["test#strategy"] = {
        nearest = "vimux",
        file = "vimux",
        suite = "vimux",
      }
      vim.g["preserve_screen"] = true
      vim.g.VimuxOrientation = "h"
      vim.g.VimuxHeight = "30"
      vim.g.VimuxCloseOnExit = true
      vim.g["test#custom_strategies"] = {
        vimux_watch = function(args)
          vim.cmd "call VimuxClearTerminalScreen()"
          vim.cmd "call VimuxClearRunnerHistory()"
          vim.cmd(string.format("call VimuxRunCommand('fd . | entr -c %s')", args))
        end,
      }
      vim.keymap.set({ "n", "v" }, "<C-c><C-c>", function()
        -- yank text into v register
        if vim.api.nvim_get_mode()["mode"] == "n" then
          vim.cmd 'normal vip"vy'
        else
          vim.cmd 'normal "vy'
        end
        -- construct command with v register as command to send
        -- vim.cmd(string.format('call VimuxRunCommand("%s")', vim.trim(vim.fn.getreg('v'))))
        vim.cmd "call VimuxRunCommand(@v)"
      end)
    end,

    keys = {
      { "<leader>rr", "<CMD>VimuxPromptCommand<CR>", { desc = "run the runsman" } },
      { "<leader>r.", "<CMD>VimuxRunLastCommand<CR>", { desc = "run the last run command" } },
      { "<leader>rc", "<CMD>VimuxClearTerminalScreen<CR>", { desc = "clear the current run terminal" } },
      { "<leader>rq", "<CMD>VimuxCloseRunner<CR>", { desc = "close the runner" } },
      { "<leader>r?", "<CMD>VimuxInspectRunner<CR>", { desc = "inspect the runner" } },
      { "<leader>r!", "<CMD>VimuxInterruptRunner<CR>", { desc = "interrupt the runner (bang'er)" } },
      { "<leader>rz", "<CMD>VimuxZoomRunner<CR>", { desc = "zoom the runner" } },
      { "<leader>tt", "<cmd>TestFile<cr>", { desc = "run test watcher" } },
      {
        "<leader>tT",
        "<cmd>TestNearest -strategy=vimux_watch<cr><cr>",
        { desc = "run test for whole file" },
      },
      { "<leader>tn", "<cmd>TestNearest<cr>", { desc = "run nearest test to cursor" } },
      {
        "<leader>tN",
        "<cmd>TestNearest -strategy=vimux_watch<cr><cr>",
        { desc = "run nearest test to cursor" },
      },
      { "<leader>ts", "<cmd>TestSuite<cr>", { desc = "run test suite" } },
      { "<leader>tS", "<cmd>TestSuite -strategy=vimux_watch<cr><cr>", { desc = "run test suite" } },
      { "<leader>t.", "<cmd>TestLast<cr>", { desc = "re-run the last test run" } },
      {
        "<leader>t>",
        "<cmd>TestLast -strategy=vimux_watch_side_split<cr><cr>",
        { desc = "re-run the last test run" },
      },
      { "<leader>tv", "<cmd>TestVisit<cr>", { desc = "visit the last run test" } },
    },
  },
}
