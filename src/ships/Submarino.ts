import { Coordinate, Orientation, Ship } from './objects'

class Submarino extends Ship {
  public constructor (start?: Coordinate, orientation?: Orientation) {
    super('Submarino', 1, start, orientation)
  }
}

export default Submarino
