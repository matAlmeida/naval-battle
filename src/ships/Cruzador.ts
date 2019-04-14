import { Coordinate, Orientation, Ship } from '../objects'

class Cruzador extends Ship {
  public constructor (start?: Coordinate, orientation?: Orientation) {
    super('Cruzador', 4, start, orientation)
  }
}

export default Cruzador
