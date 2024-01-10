import Player from './player'

// The main game state
export default class GameState {
  players: Player[];

  constructor() {
    this.players = []
  }

  addNewPlayer(name: string) {
    this.players.push(new Player(name))
  }

  getPlayers() {
    return this.players
  }
}
