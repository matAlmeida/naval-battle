import { Coordinate, Orientation, Ship } from './objects'

class Destroyer extends Ship {
  public constructor (start?: Coordinate, orientation?: Orientation) {
    super('Destroyer', 2, start, orientation)
  }
}

export default Destroyer
