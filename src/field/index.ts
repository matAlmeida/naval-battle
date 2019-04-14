class Field {
  private sea: number[][]
  private size: number

  public constructor (size: number = 10) {
    this.size = size
    this.cleanSea()
  }

  private cleanSea (): void {
    for (let y = 0; y < this.size; y++) {
      let row: number[] = []
      for (let x = 0; x < this.size; x++) {
        row.push(0)
      }
      this.sea.push(row)
    }
  }
}
export { Field }
