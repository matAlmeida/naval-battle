const SpecialField = require("../src/objects/Field").SpecialField;

describe("SpecialField", () => {
  describe("should", () => {
    describe("create 5x5 field", () => {
      const field = new SpecialField(5);
      const fieldArray = field.fieldAsArray;

      it("with size 5", () => {
        const size = field.size;

        expect(size).toBe(5);
      });

      it("with 25 items", () => {
        expect(fieldArray.length).toBe(25);
      });

      it("filled with nulls", () => {
        const hasNull = fieldArray.find(item => item !== null);

        expect(!!hasNull).toBe(false);
      });

      describe("Spaceship", () => {
        it("place at (1, 1) DOWN", () => {
          const placed = field.placeSpaceship(
            1,
            1,
            SpecialField.DIRECTIONS.DOWN,
            "ss"
          );

          expect(placed).toBe(true);
          field.cleanSpaceship(1, 1, SpecialField.DIRECTIONS.DOWN);
        });

        it("place at (1, 1) RIGHT", () => {
          const placed = field.placeSpaceship(
            1,
            1,
            SpecialField.DIRECTIONS.RIGHT,
            "ss"
          );

          expect(placed).toBe(true);
          field.cleanSpaceship(1, 1, SpecialField.DIRECTIONS.RIGHT);
        });
      });
    });
  });
});
