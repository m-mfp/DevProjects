import tkinter as tk
from tkinter import ttk
import settings
from game import Game

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


        game = Game(self.top_frame, self.left_frame, self.center_frame)
        game.startmenu()



if __name__ == "__main__":
    app = App()
    app.mainloop()