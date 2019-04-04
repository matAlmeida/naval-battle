class Game {
  private rows = 10
  private column = 10
  private maxShips = 10
  private remainingShips = 0

  private field: number[]

  public constructor () {
    this.clearField()
  }

  public clearField () : void {
    this.field = Array(this.rows * this.column).fill(0)
  }

  public showField () : void {
    this.field.map((square, index): void => {
      if ((index % this.column) === (this.column - 1)) {
        process.stdout.write(`\n`)
        return
      }
      process.stdout.write(`${square}\t`)
    })
  }

  private getRandSquare () : number[] {
    const x = parseInt((Math.random() * 100).toFixed(0)) % this.rows
    const y = parseInt((Math.random() * 100).toFixed(0)) % this.column

    return [x, y]
  }

  public setShips () : void {
    let s = 0

    while (s < this.maxShips) {
      const [x, y] = this.getRandSquare()
      if (this.field[this.rows * x + y] !== 1) {
        s += 1
        this.field[this.rows * x + y] = 1
      }
    }

    this.remainingShips = s
  }

  public leftShips () : number {
    return this.remainingShips
  }

  public atack (x: number, y: number) : boolean {
    if (this.field[this.rows * x + y] && (this.field[this.rows * x + y] === 1)) {
      this.remainingShips -= 1
      this.field[this.rows * x + y] = 2
      return true
    }

    return false
  }
}

const game = new Game()

game.showField()
console.log('----------------------')
game.setShips()
game.showField()
game.atack(3, 2)
game.atack(2, 2)
game.atack(1, 2)
console.log('Remaining ships:', game.leftShips())
console.log('----------------------')
game.showField()
