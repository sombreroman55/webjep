export interface Clue {
  question: string;
  answer: string;
  value: number;
}

export interface Category {
  title: string;
  clues: Clue[];
}

export interface Round {
  categories: Category[];
}

export interface FinalRound {
  category: string;
  question: string;
  answer: string;
}

export interface Game {
  rounds: Round[];
  finalRound: FinalRound;
}

export function newGame(): Game {
  return {
    rounds: [
      {
        categories: [
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 200,
              },
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 600,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1000,
              },
            ]
          },
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 200,
              },
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 600,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1000,
              },
            ]
          },
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 200,
              },
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 600,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1000,
              },
            ]
          },
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 200,
              },
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 600,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1000,
              },
            ]
          },
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 200,
              },
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 600,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1000,
              },
            ]
          },
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 200,
              },
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 600,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1000,
              },
            ]
          },
        ]
      },
      {
        categories: [
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1200,
              },
              {
                question: "",
                answer: "",
                value: 1600,
              },
              {
                question: "",
                answer: "",
                value: 2000,
              },
            ]
          },
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1200,
              },
              {
                question: "",
                answer: "",
                value: 1600,
              },
              {
                question: "",
                answer: "",
                value: 2000,
              },
            ]
          },
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1200,
              },
              {
                question: "",
                answer: "",
                value: 1600,
              },
              {
                question: "",
                answer: "",
                value: 2000,
              },
            ]
          },
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1200,
              },
              {
                question: "",
                answer: "",
                value: 1600,
              },
              {
                question: "",
                answer: "",
                value: 2000,
              },
            ]
          },
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1200,
              },
              {
                question: "",
                answer: "",
                value: 1600,
              },
              {
                question: "",
                answer: "",
                value: 2000,
              },
            ]
          },
          {
            title: "",
            clues: [
              {
                question: "",
                answer: "",
                value: 400,
              },
              {
                question: "",
                answer: "",
                value: 800,
              },
              {
                question: "",
                answer: "",
                value: 1200,
              },
              {
                question: "",
                answer: "",
                value: 1600,
              },
              {
                question: "",
                answer: "",
                value: 2000,
              },
            ]
          },
        ]
      },
    ],
    finalRound: {
      category: "",
      question: "",
      answer: "",
    }
  };
}
