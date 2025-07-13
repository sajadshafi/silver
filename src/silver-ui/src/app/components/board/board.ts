import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { GameService } from '../../services/game.service';
import { ChangeDetectorRef } from '@angular/core';
import { Board, SelectedPiece, GetValidMoveRequestModel } from '../../models/game.model';

@Component({
    selector: 'app-board',
    templateUrl: './board.html',
    imports: [CommonModule, RouterLink],
    styleUrls: ['./board.scss'],
})
export class BoardComponent {
    ranks = [8, 7, 6, 5, 4, 3, 2, 1];
    files = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'];

    selectedPiece: SelectedPiece | null = null;

    validMoves: number[] = [];

    board: Board = [];

    constructor(private gameService: GameService, private cdr: ChangeDetectorRef) { }

    ngOnInit(): void {
        this.startNewGame();
    }

    startNewGame(): void {
        this.gameService.initializeGame().subscribe({
            next: (game) => {
                this.board = game.board;
                this.cdr.detectChanges();
            },
            error: (err) => {
                console.error('Failed to initialize game', err);
            }
        });
    }

    getValidMoves(): void {
        const model: GetValidMoveRequestModel = {
            piece: this.selectedPiece?.piece,
            board: this.board,
            position: this.selectedPiece?.cell
        };

        this.gameService.getValidMoves(model).subscribe({
            next: (validMoves) => {
                this.validMoves = validMoves;
                this.cdr.detectChanges();
            },
            error: (err) => {
                console.error('Failed to initialize game', err);
            }
        });
    }

    onCellSelected = (index: number) => {
        if (this.board[index] && this.board[index].piece) {
            this.selectedPiece = {
                piece: this.board[index].piece,
                cell: index
            }

            this.getValidMoves();
        }
        else if (this.selectedPiece && this.isValidMoveCell(index)) {
            this.board[index].piece = this.selectedPiece.piece;
            this.board[this.selectedPiece.cell] = {};
            this.selectedPiece = null; // Deselect if clicked on an empty cell
            this.validMoves = [];
        } else {
            this.selectedPiece = null; // Deselect if clicked on an empty cell
            this.validMoves = [];
        }
    }

    isValidMoveCell(targetIndex: number): boolean {
        return this.validMoves?.some(move => move === targetIndex);
    }

    getCellBorderClass(index: number): string {
        return this.isValidMoveCell(index) ? 'border-circle-valid-move-cell' : 'board-circle';
    }

    getCellBgClass(index: number): string {
        const [rank, file] = this.getFileAndRank(index);
        if (this.selectedPiece?.cell === index) {
            return (rank + file) % 2 === 0 ? 'bg-[#97b6c2]' : 'bg-[#358096]';
        }
        return (rank + file) % 2 === 0 ? 'bg-[#d7eaf4]' : 'bg-[#83b8c8]';
    }


    getIndex(rank: number, file: number): number {
        return (8 * (rank - 1)) + (file - 1); // from white's perspective (rank 8 on top)
    }

    getFileAndRank(index: number): [number, number] {
        let rank = Math.floor(index / 8) + 1;
        let file = (index % 8) + 1;
        return [rank, file];
    }

}

