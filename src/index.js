class Field {
  constructor(size) {
    this.__size = size;
    this.field = new Array(size * size).fill(5);
  }

  get(x, y) {
    const position = this.__size * y + x;

    return this.field[position];
  }

  put(x, y, item) {
    const position = this.__size * y + x;

    try {
      this.field[position] = item;
      return true;
    } catch (error) {
      return false;
    }
  }

  get size() {
    return this.__size;
  }
}
