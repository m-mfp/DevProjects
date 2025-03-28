import tkinter as tk
from tkinter import ttk
import settings
from cell import Cell


class Game:
    def __init__(self, top_frame, left_frame, center_frame):
        self.top_frame = top_frame
        self.left_frame = left_frame
        self.center_frame = center_frame
        self.first_game = True
        Cell.restart_callback = self.start_game

        self.difficulties = (
            ('EASY', 6),
            ('MEDIUM', 8),
            ('HARD', 10),
            ('ABSURD', 12)
        )
        self.selected_difficulty = tk.StringVar(value=self.difficulties[0][1])

    def create_play_button(self):
        self._playbutton = ttk.Button(
            self.left_frame,
            text='PLAY',
            command=self.start_game,
        )
        self._playbutton.pack()
    
    def create_difficulty_button(self):        
        for widget in self.left_frame.winfo_children():
            widget.destroy()

        label = ttk.Label(self.left_frame, text="CHOOSE DIFFICULTY:")
        label.pack(padx=5, pady=5)

        for diff in self.difficulties:
            self.difficulty = ttk.Radiobutton(
                self.left_frame,
                text=diff[0],
                value=diff[1],
                variable=self.selected_difficulty,
            )
            self.difficulty.pack(padx=5, pady=3)

    def startmenu(self):
        if self.first_game:
            self.first_game = False
            self.create_difficulty_button()
            self.create_play_button()
        else:
            self.create_difficulty_button()
            self.create_play_button()
            self._playbutton.configure(text='RESTART')
            self.counter()

    def counter(self):
        Cell.create_cell_count_label(self.left_frame)
        Cell.cell_count_label_object.pack(padx=5, pady=3)


    def start_game(self):
        for widget in self.center_frame.winfo_children():
            widget.destroy()

        Cell.all = []
        settings.GRID_SIZE = int(self.selected_difficulty.get())
        settings.CELL_COUNT = settings.GRID_SIZE ** 2
        settings.MINES_COUNT = settings.CELL_COUNT // 4
        Cell.cell_count = settings.CELL_COUNT

        for x in range(settings.GRID_SIZE):
            for y in range(settings.GRID_SIZE):
                c = Cell(x, y)
                c.create_btn_object(self.center_frame)
                c.cell_btn_object.grid(column=x, row=y, sticky='snew', padx=4, pady=4)
                self.center_frame.grid_columnconfigure(x, weight=1)
                self.center_frame.grid_rowconfigure(y, weight=1)

        Cell.randomize_mines()
        self.startmenu()
        