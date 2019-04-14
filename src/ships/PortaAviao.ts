import { Coordinate, Orientation, Ship } from './objects'

class PortaAviao extends Ship {
  public constructor (start?: Coordinate, orientation?: Orientation) {
    super('Porta Avião', 5, start, orientation)
  }
}

export default PortaAviao
