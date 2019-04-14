import { Coordinate, Orientation } from '.'

class Ship {
  public name: string
  public start: Coordinate
  public orientation: Orientation
  public size: number

  public constructor (name: string, size: number, start: Coordinate = { x: 0, y: 0 }, orientation: Orientation = 'horizontal') {
    this.name = name
    this.size = size
    this.start = start
    this.orientation = orientation
  }
}

export default Ship
