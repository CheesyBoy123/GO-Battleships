package main
import ("fmt"
        "bufio"
        "os"
        "strings"
        "strconv"
        "math/rand"
        "time"
)

var our_board[8][8] rune;
var our_board_hits[8][8] rune;
var opponent_board[8][8] rune;
var opponent_board_hits[8][8] rune;

//all the different ship sizes.
var ships[5] int;

func main() {

  initBoards();
  startGame();

}

func initBoards() {
  ships = [5]int {5, 4, 3, 3, 2};

  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      our_board[i][j] = ' ';
      our_board_hits[i][j] = ' ';
      opponent_board[i][j] = ' ';
      opponent_board_hits[i][j] = ' ';
    }
  }

}



func startGame() {
  totalShipHits := 0;
  for i := 0; i < len(ships); i++ {
    totalShipHits += ships[i];
  }

  playerHits := 0;
  oppHits := 0;
  currentTurn := 0;

  pickPlayerShips();
  printBoard(our_board);
  fmt.Println();
  fmt.Println();
  pickRandomShips();

  reader := bufio.NewReader(os.Stdin);

  for {
    if(playerHits == totalShipHits || oppHits == totalShipHits) {
      break;
    }

    //our turn
    if(currentTurn % 2 == 0) {
      printBoard(our_board_hits);
      fmt.Printf("------------------------\n");
      fmt.Printf("Our Board\n");
      fmt.Printf("------------------------\n");
      fmt.Printf("Make column selection (a-h):\n");
      printBoard(our_board);
      text, _ := reader.ReadString('\n');

      text = strings.Replace(text, "\n", "", -1);
      if(text[0] < 97 || text[0] > 104) {
        fmt.Printf("Invalid column given (%c), please try again.\n", text[0]);

        continue;
      }

      //get our column # (a == 0, h == 6)
      col := int(text[0] - 97);
      fmt.Printf("Make a row selection (1-8):\n");

      row_text, _ := reader.ReadString('\n');
      row_text = strings.Replace(row_text, "\n", "", -1);

      row, err := strconv.Atoi(row_text);
      row--;
      if(err != nil) {
          fmt.Printf("Row not found, please try again.\n");

          continue;
      }

      if(row < 0 || row > 8) {
        fmt.Printf("Invalid row (%d) given, please try again.\n", row +1);

        continue;
      }

      if(our_board_hits[row][col] != ' ') {

        fmt.Printf("You've already selected that cord, try a different one!\n");
        continue;
      }
      //HIT!!!
      if(opponent_board[row][col] == 'S') {
        fmt.Printf("HIT!\n");
        our_board_hits[row][col] = 'H';
        playerHits++;
      } else {
        fmt.Printf("MISS!\n");
        our_board_hits[row][col] = 'M';
      }
    //AI's turn
    } else {
      s1 := rand.NewSource(time.Now().UnixNano());
      r1 := rand.New(s1);
      x := r1.Intn(8);
      y := r1.Intn(8);

      if(opponent_board_hits[x][y] != ' ') {
        continue;
      }

      fmt.Println();
      if(our_board[x][y] == 'S') {
        oppHits++;
        fmt.Printf("~~~Opponent hit our ship!~~~\n");
        our_board[x][y] = 'H'
      } else {
        fmt.Printf("Opponent missed our ship!\n");
      }
      fmt.Println();

    }
    currentTurn++;


  }

  if(playerHits == totalShipHits) {
    fmt.Printf("Congrats you beat our super sophisticated AI!!!\n");
  } else {
    fmt.Printf("Better luck next time our AI is super hard to beat!\n");
  }

}

func pickRandomShips() {

  for i := 0; i < len(ships); i++ {
    s1 := rand.NewSource(time.Now().UnixNano());
    r1 := rand.New(s1);

    x := r1.Intn(8);
    y := r1.Intn(8);

    dir := r1.Intn(4);
    //N S E W
    if(dir == 0) {
      if(x - ships[i] < 0 || !checkValid(opponent_board, x, y, dir -1 ,0,ships[i])) {
        i--;
        continue;
      }

      for j := x; j > x-ships[i]; j-- {
          opponent_board[j][y] = 'S';
      }

    }else if(dir == 1) {
      if(x + ships[i] >= 8 || !checkValid(opponent_board, x, y, 1 ,0,ships[i])) {
        i--;
        continue;
      }
      for j := x; j < x+ships[i]; j++ {
          opponent_board[j][y] = 'S';
      }

    }else if(dir == 2) {
      if( y + ships[i] >=8 || !checkValid(opponent_board, x, y,  0 ,1,ships[i])) {
        i--;
        continue;
      }

      for j := y; j < y+ships[i]; j++ {
          opponent_board[x][j] = 'S';
      }

    }else if(dir == 3) {
      if(y - ships[i] < 0 || !checkValid(opponent_board, x, y,  0 , -1, ships[i])) {
        i--;
        continue;
      }
      for j := y; j > y-ships[i]; j-- {
          opponent_board[x][j] = 'S';
      }
    }
  }
}

func pickPlayerShips() {

  reader := bufio.NewReader(os.Stdin);

  for i := 0; i < len(ships); i++ {
    printBoard(our_board);
    fmt.Printf("Time to pick a location for the ship of length: %d\n", ships[i]);
    fmt.Printf("Please Enter a column character (a-f):\n");

    text, _ := reader.ReadString('\n');

    text = strings.Replace(text, "\n", "", -1);
    if(text[0] < 97 || text[0] > 104) {
      fmt.Printf("Invalid column given (%c), please try again.\n", text[0]);
      i--;
      continue;
    }

    //get our column # (a == 0, h == 6)
    col := int(text[0] - 97);

    fmt.Printf("Please enter a row number (1-8):\n");

    row_text, _ := reader.ReadString('\n');
    row_text = strings.Replace(row_text, "\n", "", -1);

    row, err := strconv.Atoi(row_text);
    row--;
    if(err != nil) {
        fmt.Printf("Row not found, please try again.\n");
        i--;
        continue;
    }

    if(row < 0 || row > 8) {
      fmt.Printf("Invalid row (%d) given, please try again.\n", row +1);
      i--;
      continue;
    }

    fmt.Printf("Please enter a ship orientation (n,s,e,w):\n");
    orient, _ := reader.ReadString('\n');
    orient = strings.Replace(orient, "\n", "", -1);


    if(orient == "n" || orient == "N") {
      if(row - ships[i] < 0) {
        fmt.Printf("Invalid placement of ship (goes off the board)!\n")
        i--;
        continue;
      }

      if(!checkValid(our_board, row, int(col), -1, 0, ships[i])) {
        fmt.Printf("Invalid placement of ship (collides with another ship)!\n");
        i--;
        continue;
      }

      for j := row; j > row-ships[i]; j-- {
          our_board[j][col] = 'S';
      }

    } else if(orient == "s" || orient == "S") {
      if(row + ships[i] >= 8) {
        fmt.Printf("Invalid placement of ship (goes off the board)!\n")
        i--;
        continue;
      }

      if(!checkValid(our_board, row, int(col), 1, 0, ships[i])) {
        fmt.Printf("Invalid placement of ship (collides with another ship)!\n");
        i--;
        continue;
      }
      for j := row; j < row+ships[i]; j++ {
          our_board[j][col] = 'S';
      }

    } else if(orient == "e" || orient == "E") {
      if(col + ships[i] >= 8) {
        fmt.Printf("Invalid placement of ship (goes off the board)!\n")
        i--;
        continue;
      }

      if(!checkValid(our_board, row, int(col), 0, 1, ships[i])) {
        fmt.Printf("Invalid placement of ship (collides with another ship)!\n");
        printBoard(our_board);
        i--;
        continue;
      }

      for j := col; j < col+ships[i]; j++ {
          our_board[row][j] = 'S';
      }
    } else if(orient == "w" || orient == "W") {
      if(col - ships[i] < 0) {
        fmt.Printf("Invalid placement of ship (goes off the board)!\n")
        i--;
        continue;
      }

      if(!checkValid(our_board, row, int(col), 0, -1, ships[i])) {
        fmt.Printf("Invalid placement of ship (collides with another ship)!\n");
        i--;
        continue;
      }

      for j := col; j > col-ships[i]; j-- {
          our_board[row][j] = 'S';
      }

    } else {
      fmt.Printf("Orientation not found, please try again.\n");
      i--;
      continue;
    }

  }

}

func printBoard(board[8][8] rune) {

  fmt.Printf("  | a | b | c | d | e | f | g | h |\n");
  for i := 0; i < 8; i++ {
    fmt.Printf("%d ", i +1);
    for j := 0; j < 8; j++ {
      fmt.Printf("| %c ", board[i][j]);
    }
    fmt.Printf("|\n");

  }
}

func checkValid(board[8][8] rune, start_row int, start_col int, dir_x int, dir_y int, ship_size int) bool {
  if(dir_x != 0) {
    //check the row to see if it's valid.
    for i := start_row; i != start_row + (ship_size * dir_x); i += dir_x {
      if(board[i][start_col] != ' ') {
        return false;
      }
    }
    return true;
  } else if(dir_y != 0) {
    for j := start_col; j != start_col + (ship_size * dir_y); j += dir_y {
      if(board[start_row][j] != ' ') {
        fmt.Printf("%d %d \n", start_row, j);
        return false;
      }
    }

    return true;
  } else {
    return false;
  }
}
