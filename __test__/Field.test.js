const Field = require("../src/objects/Field");

describe("Field", () => {
  describe("should", () => {
    describe("create 5x5 field", () => {
      const field = new Field(5);
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

      describe("with 0 in (1, 1)", () => {
        it("place", () => {
          const placed = field.place(1, 1, 0);

          expect(placed).toBe(true);
        });

        it("was placed", () => {
          const item = field.get(1, 1);

          expect(item).toBe(0);

          field.cleanPosition(1, 1);
        });
      });

      describe("with 0 in (1, 2)", () => {
        it("place", () => {
          const placed = field.place(1, 2, 0);

          expect(placed).toBe(true);
        });

        it("was placed", () => {
          const item = field.get(1, 2);

          expect(item).toBe(0);

          field.cleanPosition(1, 2);
        });
      });

      describe("with 0 in (1, 1) and (1, 2)", () => {
        it("place", () => {
          const aplaced = field.place(1, 1, 0);
          const bplaced = field.place(1, 2, 0);

          expect(aplaced && bplaced).toBe(true);
        });

        it("was placed", () => {
          const aitem = field.get(1, 1);
          const bitem = field.get(1, 2);

          expect(aitem).toBe(0);
          expect(bitem).toBe(0);

          field.cleanPosition(1, 1);
          field.cleanPosition(1, 2);
        });
      });

      describe("with a set of item", () => {
        const ship = ["s1", "s1", "s1"];

        it("place DOWN", () => {
          const placed = field.placeSet(1, 1, Field.DIRECTIONS.DOWN, ship);

          expect(placed).toBe(true);
        });

        it("was placed DOWN", () => {
          const placedCorrectly = !ship.find(
            (item, index) => field.get(1, 1 + index) !== item
          );

          expect(placedCorrectly).toBe(true);

          field.cleanSet(1, 1, Field.DIRECTIONS.DOWN, ship.length);
        });

        it("place RIGHT", () => {
          const placed = field.placeSet(1, 1, Field.DIRECTIONS.RIGHT, ship);

          expect(placed).toBe(true);
        });

        it("was placed RIGHT", () => {
          const placedCorrectly = !ship.find(
            (item, index) => field.get(1 + index, 1) !== item
          );

          expect(placedCorrectly).toBe(true);

          field.cleanSet(1, 1, Field.DIRECTIONS.RIGHT, ship.length);
        });
      });
    });
  });
});
