import { Routes } from '@angular/router';
import { BoardComponent } from './components/board/board';
import { HomeComponent } from './components/home/home';

export const routes: Routes = [
  { path: '', component: HomeComponent },
  { path: 'board', component: BoardComponent },
];
