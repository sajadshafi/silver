import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Game, GetValidMoveRequestModel } from '../models/game.model';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class GameService {
    private baseUrl = 'http://localhost:8080';

    constructor(private http: HttpClient) { }

    initializeGame(): Observable<Game> {
        return this.http.get<Game>(`${this.baseUrl}/game/initialize`);
    }

    getValidMoves(model: GetValidMoveRequestModel): Observable<number[]> {
        return this.http.post<number[]>(`${this.baseUrl}/game/get-valid-moves`, model);
    }

}

