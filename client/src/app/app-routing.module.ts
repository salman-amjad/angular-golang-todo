import { NgModule }             		from '@angular/core';
import { RouterModule, Routes, Data } 	from '@angular/router';
import { TodoListComponent } from './todo-list/todo-list.component';
import { TodoFormComponent } from './todo-form/todo-form.component';

const site_title: string = 'Salman Amjad';
const sep: string = ' - ';

const routes: Routes = [
	{ 
		path: '',  			
		component: TodoListComponent
	},
  	{ 
		path: 'add', 		
		component: TodoFormComponent
	}
];
@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule {}
