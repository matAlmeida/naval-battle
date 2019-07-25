const Field = require("./Field");

class SpecialField extends Field {
  constructor(size) {
    super(size);
  }

  placeSpaceship(x, y, direction, item) {
    if (direction === this.__DIRECTIONS.DOWN) {
      const top = this.get(x, y) === null;
      const left = this.get(x - 1, y + 1) === null;
      const right = this.get(x + 1, y + 1) === null;

      if (top && left && right) {
        this.place(x, y, item);
        this.place(x - 1, y + 1, item);
        this.place(x + 1, y + 1, item);

        return true;
      }
    } else if (direction === this.__DIRECTIONS.RIGHT) {
      const top = this.get(x, y) === null;
      const left = this.get(x + 1, y + 1) === null;
      const right = this.get(x + 1, y - 1) === null;

      if (top && left && right) {
        this.place(x, y, item);
        this.place(x + 1, y + 1, item);
        this.place(x + 1, y - 1, item);

        return true;
      }
    } else if (direction === this.__DIRECTIONS.UP) {
      const top = this.get(x, y) === null;
      const left = this.get(x + 1, y - 1) === null;
      const right = this.get(x - 1, y - 1) === null;

      if (top && left && right) {
        this.place(x, y, item);
        this.place(x + 1, y - 1, item);
        this.place(x - 1, y - 1, item);

        return true;
      }
    } else if (direction === this.__DIRECTIONS.LEFT) {
      const top = this.get(x, y) === null;
      const left = this.get(x - 1, y - 1) === null;
      const right = this.get(x - 1, y + 1) === null;

      if (top && left && right) {
        this.place(x, y, item);
        this.place(x - 1, y - 1, item);
        this.place(x - 1, y + 1, item);

        return true;
      }
    }

    return false;
  }

  cleanSpaceship(x, y, direction) {
    if (direction === this.__DIRECTIONS.DOWN) {
      this.cleanPosition(x, y);
      this.cleanPosition(x - 1, y + 1);
      this.cleanPosition(x + 1, y + 1);
    } else if (direction === this.__DIRECTIONS.DOWN) {
      this.cleanPosition(x, y);
      this.cleanPosition(x + 1, y + 1);
      this.cleanPosition(x + 1, y - 1);
    }
  }
}

module.exports = SpecialField;
