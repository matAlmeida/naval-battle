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
    if (x >= 0 && x < this.__size && y >= 0 && y < this.__size) {
      const position = this.__size * y + x;

      return this.__field[position];
    }

    return undefined;
  }

  get size() {
    const size = 0 + this.__size;

    return size;
  }

  get fieldAsArray() {
    return [...this.__field];
  }

  toString() {
    let string = ""

    for(let i = 0; i < this.__size; i++) {
      for(let j = 0; j < this.__size; j++) {
        string += `${this.get(i, j)}` + "\t"
      }
      string += "\n"
    }

    return string
  }

  checkIfCanPlace(x, y, item, hasSame = "") {
    const p = this.get(x, y);
    const up = this.get(x, y - 1);
    const right = this.get(x + 1, y);
    const down = this.get(x, y + 1);
    const left = this.get(x - 1, y);

    const noAdjacent = !up && !right && !down && !left;
    let same = false;

    if (hasSame === "") {
      same = up === item || right === item || down === item || left === item;
    } else {
      same =
        up === item[hasSame] ||
        right === item[hasSame] ||
        down === item[hasSame] ||
        left === item[hasSame];
    }

    if (p === null && (noAdjacent || same)) {
      return true;
    }

    return false;
  }

  place(x, y, item, idField = "") {
    const position = this.__size * y + x;

    try {
      const canPlace = this.checkIfCanPlace(x, y, item, idField);
      if (canPlace === true) {
        this.__field[position] = item;
        return true;
      }

      return false;
    } catch (error) {
      return false;
    }
  }

  cleanPosition(x, y) {
    if (x >= 0 && x < this.__size && y >= 0 && y < this.__size) {
      const position = this.__size * y + x;
      this.__field[position] = null;

      return true;
    }

    return false;
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

  cleanSet(x, y, direction, lenght) {
    try {
      if (direction === DIRECTIONS.DOWN) {
        for (let i = 0; i < lenght; i++) {
          this.cleanPosition(x, y + i);
        }

        return true;
      } else if (direction === DIRECTIONS.RIGHT) {
        for (let i = 0; i < lenght; i++) {
          this.cleanPosition(x + i, y);
        }

        return true;
      }

      return false;
    } catch (error) {
      return false;
    }
  }
}

module.exports = Field;
