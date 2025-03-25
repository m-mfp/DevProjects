from tkinter import Button, Label, messagebox
import random, settings
import sys

class Cell:
    all = []
    cell_count = settings.CELL_COUNT
    cell_count_label_object = None

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
            width = 12,
            height = 4,
            bg = 'ivory'
        )

        # Configure what happens when clicking the button
        btn.bind('<Button-1>', self.left_click_actions)
        btn.bind('<Button-3>', self.right_click_actions)
        self.cell_btn_object = btn

    @staticmethod
    def create_cell_count_label(location):
        lbl = Label(
            location,
            bg = 'black',
            fg = 'white',
            text = f"{Cell.cell_count}",
            width = 12,
            height = 4,
            font = ('', 30),
        )
        Cell.cell_count_label_object = lbl

    def left_click_actions(self, event):
        if self.is_mine:
            self.show_mine()
        else:
            if self.surrounding_mines_count == 0:
                for cell_obj in self.surrounded_cells:
                    cell_obj.show_cell()
            self.show_cell()

            if Cell.cell_count == settings.MINES_COUNT:
                messagebox.showinfo("Game Over", "You Won!")

        self.cell_btn_object.unbind('<Button-1>')
        self.cell_btn_object.unbind('<Button-3>')

    def get_cell_by_axis(self, x, y):
        for cell in Cell.all:
            if cell.x == x and cell.y == y:
                return cell

    @property
    def surrounded_cells(self):
        cells = [
            self.get_cell_by_axis(self.x - 1, self.y - 1),
            self.get_cell_by_axis(self.x - 1, self.y),
            self.get_cell_by_axis(self.x - 1, self.y + 1),
            self.get_cell_by_axis(self.x, self.y - 1),
            self.get_cell_by_axis(self.x, self.y + 1),
            self.get_cell_by_axis(self.x + 1, self.y - 1),
            self.get_cell_by_axis(self.x + 1, self.y),
            self.get_cell_by_axis(self.x + 1, self.y + 1),
        ]
        cells = [cell for cell in cells if cell != None]
        return cells

    @property
    def surrounding_mines_count(self):
        counter = 0
        for cell in self.surrounded_cells:
            if cell.is_mine:
                counter += 1
        return counter

    def show_cell(self):
        # NEEDS TO BE FIXED - IF NEXT CELL IS ALSO ZERO IT SHOULD BE OPENED TOO, BUT IT IS JUST BEIGN SHOWN
        if not self.is_open:
            Cell.cell_count -= 1
            self.cell_btn_object.configure(
                text=self.surrounding_mines_count, 
            )
            if Cell.cell_count_label_object:
                Cell.cell_count_label_object.configure(text=f"{Cell.cell_count}")
        
        self.cell_btn_object.configure(
            bg='ivory'
        )

        self.is_open = True

    def show_mine(self):
        self.cell_btn_object.configure(bg='red')
        retry = messagebox.askyesno("Game Over", "You clicked on a mine!\nDo you want to retry?")
        if retry:
            self.restart_game()
        else:
            pass#sys.exit()

    def right_click_actions(self, event):
        if not self.is_mine_candidate:
            self.cell_btn_object.configure(
                bg='orange'
            )
            self.is_mine_candidate = True
        else:
            self.cell_btn_object.configure(
                bg='ivory'    # SystemButtonFace didn't work
            )
            self.is_mine_candidate = False

    def restart_game(self):
        pass#sys.exit()

    @staticmethod
    def randomize_mines():
        picked_cells = random.sample(Cell.all, settings.MINES_COUNT)
        for picked_cell in picked_cells:
            picked_cell.is_mine = True

    def __repr__(self):
        return f"Cell({self.x}, {self.y})"