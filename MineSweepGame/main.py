import tkinter as tk
from tkinter import ttk
import settings, utils
from cell import Cell

class App(tk.Tk):
    def __init__(self):
        super().__init__()
        self.configure(bg="black")
        self.geometry(f'{settings.WIDTH}x{settings.HEIGHT}')
        self.title("MineSweep Game")

        self.top_frame = tk.Frame(
            self,
            bg = "black",
            width = settings.WIDTH,
            height = utils.height_prct(25),
        )
        self.top_frame.place(x=0, y=0)

        self.left_frame = tk.Frame(
            self,
            bg = "black",
            width = utils.width_prct(25),
            height = utils.height_prct(75),
        )
        self.left_frame.place(x=0, y=utils.height_prct(25))

        self.center_frame = tk.Frame(
            self,
            bg = "black",
            width = utils.width_prct(75),
            height = utils.height_prct(75),
        )
        self.center_frame.place(x=utils.width_prct(25), y=utils.height_prct(25))

        self.start_menu()    

    def start_menu(self):
        self.playbutton = tk.Button(
            self.top_frame,
            text='PLAY',
            bg='grey',
            fg='ivory',
            command=self.start_game,
        )
        self.playbutton.pack(expand=True)

    def start_game(self):
        for x in range(settings.GRID_SIZE):
            for y in range(settings.GRID_SIZE):
                c = Cell(x, y)
                c.create_btn_object(self.center_frame)
                c.cell_btn_object.grid(column=x, row=y)

        Cell.create_cell_count_label(self.left_frame)
        Cell.cell_count_label_object.place(x=0, y=0)
        Cell.randomize_mines()

if __name__ == "__main__":
    app = App()
    app.mainloop()