export default class Player {
  name: string;
  score: number;

  constructor(name: string) {
    this.name = name;
    this.score = 0;
  }

  increaseScore(value: number) {
    this.score += value;
  }

  decreaseScore(value: number) {
    this.score -= value;
  }
}
