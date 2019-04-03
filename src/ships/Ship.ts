interface Coordinate {
  x: number,
  y: number
}

type Side = 'left' | 'top' | 'right' | 'bottom'

class Ship {
  public size: number
  public name: string
  public startCoordinate: Coordinate
  public frontSide: Side

  public constructor (size: number, name: string) {
    this.size = size
    this.name = name
  }

  public positionateShip (startCoordinate: Coordinate, frontSide: Side) : void {
    this.startCoordinate = startCoordinate
    this.frontSide = frontSide
  }
}

export default Ship
