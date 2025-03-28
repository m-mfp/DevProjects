from tkinter import Button, Label, messagebox
import random, settings
import sys

class Cell:
    all = []
    cell_count = 0

    def __init__(self, x:int, y:int, is_mine=False):
        self.is_mine = is_mine
        self.is_open = False
        self.is_mine_candidate = False
        self.cell_btn_object = None
        self.x = x
        self.y = y

        Cell.all.append(self)

    def create_btn_object(self, location):
        btn = Button(
            location,
            bg=settings.CELL_DEFAULT_BG,
            activebackground=settings.CELL_HOVER_BG
        )

        btn.bind('<Button-1>', self.left_click_actions)
        btn.bind('<Button-3>', self.right_click_actions)
        self.cell_btn_object = btn

    @staticmethod
    def create_cell_count_label(location):
        if hasattr(Cell, 'cell_count_label_object'):
            Cell.cell_count_label_object.destroy()
        
        lbl = Label(
            location,
            text = f"\nCells left:\n\n{Cell.cell_count}",
            width = 12,
            height = 4,
            font = ('', 30),
        )
        Cell.cell_count_label_object = lbl

    def left_click_actions(self, event):
        if self.is_mine:
            self.show_mine()
        else:
            self.show_cell()
            if Cell.cell_count == settings.MINES_COUNT:
                messagebox.showinfo("Game Over", "You Won!")

        if self.cell_btn_object and self.cell_btn_object.winfo_exists():
            self.cell_btn_object.unbind('<Button-1>')
            self.cell_btn_object.unbind('<Button-3>')

    @staticmethod
    def get_cell_by_axis(x, y):
        for cell in Cell.all:
            if cell.x == x and cell.y == y:
                return cell

    @property
    def surrounded_cells(self):
        cells = [
            Cell.get_cell_by_axis(self.x - 1, self.y - 1),
            Cell.get_cell_by_axis(self.x - 1, self.y),
            Cell.get_cell_by_axis(self.x - 1, self.y + 1),
            Cell.get_cell_by_axis(self.x, self.y - 1),
            Cell.get_cell_by_axis(self.x, self.y + 1),
            Cell.get_cell_by_axis(self.x + 1, self.y - 1),
            Cell.get_cell_by_axis(self.x + 1, self.y),
            Cell.get_cell_by_axis(self.x + 1, self.y + 1),
        ]
        cells = [cell for cell in cells if cell is not None]
        return cells

    @property
    def surrounding_mines_count(self):
        counter = 0
        for cell in self.surrounded_cells:
            if cell.is_mine:
                counter += 1
        return counter

    def show_cell(self):
        if not self.is_open:
            Cell.cell_count -= 1
            self.cell_btn_object.configure(
                text=self.surrounding_mines_count 
            )
            if Cell.cell_count_label_object:
                Cell.cell_count_label_object.configure(text=f"\nCells left:\n\n{Cell.cell_count}")        
        
        self.cell_btn_object.configure(
            bg=settings.CELL_DEFAULT_BG
        )

        self.is_open = True

        if self.surrounding_mines_count == 0:
            for cell in self.surrounded_cells:
                if not cell.is_open:
                    cell.show_cell()

    def show_mine(self):
        self.cell_btn_object.configure(bg=settings.CELL_MINE_BG)
        retry = messagebox.askyesno("Game Over", "You clicked on a mine!\nDo you want to retry?")
        if retry:
            self.restart_game()
        else:
            sys.exit()

    def right_click_actions(self, event):
        if not self.is_mine_candidate:
            self.cell_btn_object.configure(
                bg=settings.CELL_FLAG_BG,
                activebackground=settings.CELL_FLAG_BG
            )
            self.is_mine_candidate = True
        else:
            self.cell_btn_object.configure(
                bg=settings.CELL_DEFAULT_BG,
                activebackground=settings.CELL_HOVER_BG
            )
            self.is_mine_candidate = False

    @staticmethod
    def restart_game():
        if Cell.restart_callback:
            Cell.restart_callback()

    @staticmethod
    def randomize_mines():
        picked_cells = random.sample(Cell.all, settings.MINES_COUNT)
        for picked_cell in picked_cells:
            picked_cell.is_mine = True

    def __repr__(self):
        return f"Cell({self.x}, {self.y})"