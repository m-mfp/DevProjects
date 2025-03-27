import tkinter as tk
from tkinter import ttk
import settings, utils
from cell import Cell

class App(tk.Tk):
    def __init__(self):
        super().__init__()
        self.configure(bg="#A9A9A9")
        self.geometry(f'{settings.WIDTH}x{settings.HEIGHT}')
        self.title("MineSweep Game")

        self.style = ttk.Style(self)
        self.style.theme_use('classic')
        
        self.grid_rowconfigure(0, weight=1)
        self.grid_rowconfigure(1, weight=1)
        self.grid_rowconfigure(2, weight=3)
        self.grid_columnconfigure(0, weight=1)
        self.grid_columnconfigure(1, weight=4)

        self.top_frame = tk.Frame(
            self,
        )
        self.top_frame.grid(row=0, column=1, columnspan=2, sticky='snew', padx=4, pady=4)
        self.top_frame.grid_propagate(False)

        self.left_frame = tk.Frame(
            self,
        )
        self.left_frame.grid(row=0, column=0, rowspan=3, sticky='snew', padx=4, pady=4)
        self.left_frame.grid_propagate(False)

        self.center_frame = tk.Frame(
            self,
        )
        self.center_frame.grid(row=1, column=1, rowspan=2, sticky='snew', padx=4, pady=4)
        self.center_frame.grid_propagate(False)


        self.start_menu()
        

    def start_menu(self):

        # PLAY BUTTON UGLY BUT WORKING
        self.playbutton = ttk.Button(
            self.left_frame,
            text='PLAY',
            command=self.start_game,
        )

        # DIFFICULTY SETTINGS radio button
        difficulties = (
            ('EASY', 6),
            ('MEDIUM', 8),
            ('HARD', 10),
            ('ABSURD', 12)
        )
        self.selected_difficulty = tk.StringVar(value=difficulties[0][1])
        
        label = ttk.Label(self.left_frame, text="CHOOSE DIFFICULTY:")
        label.pack(padx=5, pady=5)
        
        for diff in difficulties:
            self.difficulty = ttk.Radiobutton(
                self.left_frame,
                text=diff[0],
                value=diff[1],
                variable=self.selected_difficulty,
            )
            self.difficulty.pack(padx=5, pady=3)
        self.playbutton.pack()



        # ADD SOUND & MUTE button
        # SIZE dropdown button


    def start_game(self):
        settings.GRID_SIZE = int(self.selected_difficulty.get())

        for x in range(settings.GRID_SIZE):
            for y in range(settings.GRID_SIZE):
                c = Cell(x, y)
                c.create_btn_object(self.center_frame)
                c.cell_btn_object.grid(column=x, row=y, sticky='snew', padx=4, pady=4)
                self.center_frame.grid_columnconfigure(x, weight=1)
                self.center_frame.grid_rowconfigure(y, weight=1)

        Cell.create_cell_count_label(self.left_frame)
        Cell.cell_count_label_object.place(x=0, y=0)
        Cell.randomize_mines()





if __name__ == "__main__":
    app = App()
    app.mainloop()