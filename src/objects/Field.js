const DIRECTIONS = { RIGHT: "up", DOWN: "down" };

class Field {
  constructor(size) {
    this.__size = size;
    this.__field = new Array(size * size).fill(null);
    this.__DIRECTIONS = DIRECTIONS;
  }

  static get DIRECTIONS() {
    return DIRECTIONS;
  }

  get(x, y) {
    const position = this.__size * y + x;

    return this.__field[position];
  }

  place(x, y, item) {
    const position = this.__size * y + x;

    try {
      const canPlace = this.get(x, y);
      if (canPlace === null) {
        this.__field[position] = item;
        return true;
      }

      return false;
    } catch (error) {
      return false;
    }
  }

  cleanPosition(x, y) {
    const position = this.__size * y + x;
    this.__field[position] = null;

    return true;
  }

  placeSet(x, y, direction, set) {
    const places = set.reduce((agg, item, index) => {
      if (direction === DIRECTIONS.DOWN) {
        const gotItem = this.get(x, y + index);

        if (gotItem === null) {
          return [...agg, [x, y + index, item]];
        }
      } else if (direction === DIRECTIONS.RIGHT) {
        const gotItem = this.get(x + index, y);

        if (gotItem === null) {
          return [...agg, [x + index, y, item]];
        }
      }

      return [...agg, false];
    }, []);

    const canBePlaced = places.find(item => item === false) === undefined;

    if (!canBePlaced) {
      return false;
    }

    places.map(item => this.place(...item));

    return true;
  }

  get size() {
    const size = 0 + this.__size;

    return size;
  }

  get fieldAsArray() {
    return [...this.__field];
  }
}

module.exports = Field;
