export const PieceTypes = {
  Pawn: 'pawn',
  Rook: 'rook',
  Knight: 'knight',
  Bishop: 'bishop',
  Queen: 'queen',
  King: 'king',
} as const;

export const Colors = {
  White: 'white',
  Black: 'black',
} as const;

// Types
export type PieceType = typeof PieceTypes[keyof typeof PieceTypes];
export type Color = typeof Colors[keyof typeof Colors];

export interface Piece {
  type: PieceType;
  color: Color;
}

export interface Square {
  piece?: Piece;
}

export type Board = Square[];

export interface Game {
  gameId: string;
  board: Board;
}


export interface SelectedPiece {
  piece: Piece;
  cell: number;
}

export interface GetValidMoveRequestModel {
  piece?: Piece;
  board?: Board;
  position?: number;
}
