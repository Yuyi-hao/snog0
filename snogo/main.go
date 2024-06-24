package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
	"strconv"
	"math/rand"
	"time"
)


func move_snake(snake [][2]int32, direction string){
	for i := len(snake)-1; i > 0; i-- {
		snake[i] = snake[i-1]
	}
	switch dir := direction; dir{
		case "S":
		snake[0][1]++
		case "N":
		snake[0][1]--
		case "L":
		snake[0][0]--
		case "R":
		snake[0][0]++
		default:
		fmt.Println("ERROR: Invalid direction")	
	}
	
}

func handle_keypresses(snake_direc *string){
	if (rl.IsKeyDown(rl.KeyRight) && *snake_direc != "L"){
		*snake_direc = "R"
	}
	if (rl.IsKeyDown(rl.KeyLeft) && *snake_direc != "R" ){
		*snake_direc = "L"
	}
	if (rl.IsKeyDown(rl.KeyUp) && *snake_direc != "S"){
		*snake_direc = "N"
	}
	if (rl.IsKeyDown(rl.KeyDown) && *snake_direc != "N") {
		*snake_direc = "S"
	}
		
}

func drawSnake(snake [][2]int32, gameX float32, gameY float32, cell_size int32){
	for snake_cell :=1; snake_cell < len(snake); snake_cell++{
		rl.DrawRectangle(int32(gameX)+snake[snake_cell][0]*cell_size, int32(gameY)+snake[snake_cell][1]*cell_size, cell_size, cell_size, rl.Blue)
	}
	rl.DrawRectangle(int32(gameX)+snake[0][0]*cell_size, int32(gameY)+snake[0][1]*cell_size, cell_size, cell_size, rl.Red)

}

func generate_apple(head [2]int32, max_row int32, max_col int32) [2]int32{
	apple := [2]int32{int32(rand.Intn(int(max_col)))%max_col, int32(rand.Intn(int(max_row)))%max_row}
	for apple[0] == head[0] && apple[1]==head[1]{
		apple = [2]int32{int32(rand.Intn(int(max_row)))%max_row, int32(rand.Intn(int(max_row)))%max_row}
	}
	return apple

}

func hit_boundary(max_row int32, max_col int32, snake [][2]int32) bool{
	for i := 1; i < len(snake); i++ {
		if snake[i][0] == snake[0][0] && snake[i][1] == snake[0][1]{
			return true
		}
	}
	return snake[0][0] == max_col || snake[0][1] ==  max_row || snake[0][0] <  0 || snake[0][1] <  0
}

func eat_apple(head [2]int32, apple [2]int32, score *int) bool{
	if head[0] == apple[0] && head[1] == apple[1]{
		if *score%5==0 &&  *score%10==5{
			*score += 5
		}else {
			*score += 1
		}
		return true
	}
	return false
}

func inc_snake_size(snake [][2]int32){
	n := len(snake)
	snake = append(snake, snake[n-1])
}

func main() {
	screenWidth := int32(1000)
	screenHeight := int32(750)
	gameScreenWidth := float32(screenWidth)*0.76
	gameScreenHeight := float32(screenHeight)*0.86
	var gameX float32 = 10
	var gameY float32 = (float32(screenHeight) - gameScreenHeight)/2
	var cell_size int32 = 20
	var game_row = int32(gameScreenHeight)/cell_size
	var game_col = int32(gameScreenWidth)/cell_size
	head_row := game_row/2
	head_col := game_col/2
	snake := [][2]int32{
		{head_row, head_col},
		{head_row, head_col+1},
		{head_row, head_col+2},
		{head_row, head_col+3}}
	playing := false
	score := 0
	start := time.Now()
	var duration time.Duration = 200000000
	bg_img_path := "./assets/snake.png"
	base_bg_img_path := "./assets/base_bg_img.png"
	
	rl.InitWindow(screenWidth, screenHeight, "Snake Eat apple")
	bg_img := rl.LoadImage(bg_img_path)
	base_bg_img := rl.LoadImage(base_bg_img_path)
	bg_texture := rl.LoadTextureFromImage(bg_img)
	base_bg_texture := rl.LoadTextureFromImage(base_bg_img)
	rl.UnloadImage(bg_img)
	rl.UnloadImage(base_bg_img)
	bg_position := rl.NewVector2(float32(screenWidth)-float32(bg_texture.Width)*1.2, float32(screenHeight)-float32(bg_texture.Height)*1.2)
	base_bg_position := rl.NewVector2(0, 0)
	fmt.Println(bg_texture.Width, bg_texture.Height)
	rl.SetTargetFPS(60)
	snake_direc := "N"
	apple := generate_apple(snake[0], game_row, game_col)
	
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawTextureEx(base_bg_texture, base_bg_position, 0, 1, rl.White)
		rl.DrawTextureEx(bg_texture, bg_position, 0, 1.2, rl.White)

		
		// draw basic board
		rl.DrawRectangleRounded(rl.NewRectangle(gameX, gameY, gameScreenWidth, gameScreenHeight), .03, 10, rl.Green)
		drawSnake(snake, gameX, gameY, cell_size)
		rl.DrawRectangle(int32(gameX)+apple[0]*cell_size, int32(gameY)+apple[1]*cell_size, cell_size, cell_size, rl.Orange)
		rl.DrawText("Score: "+strconv.Itoa(score), 10, 20, 20, rl.Pink)
		handle_keypresses(&snake_direc)
		if playing{
			t := time.Now()
			elapsed := t.Sub(start)
			if elapsed >= duration{
				start = time.Now()
				move_snake(snake, snake_direc)
			}

		}
		if rl.IsKeyDown(rl.KeySpace){
			playing = !playing
		}
		if rl.IsKeyDown(rl.KeyR){
			snake = [][2]int32{
				{head_row, head_col},
				{head_row, head_col+1},
				{head_row, head_col+2},
				{head_row, head_col+3}}
			snake_direc = "N"
			score = 0
		}
		if eat_apple(snake[0], apple, &score){
			apple = generate_apple(snake[0], game_row, game_col)
			snake = append(snake, [2]int32{-20, -20})
		}
		
		if hit_boundary(game_row, game_col, snake){
			playing = false
		}
		rl.DrawText("Press <Space> to run game. Press <r> to restart game", 10, screenHeight-50, 20, rl.Pink)
		rl.EndDrawing()
	}
	rl.UnloadTexture(bg_texture)
	rl.CloseWindow()
}
