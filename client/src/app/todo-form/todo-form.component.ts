import { Component, OnInit } from '@angular/core';
import { Location } from '@angular/common';
import { HttpClient, HttpHeaders, HttpRequest } from '@angular/common/http';

interface Todo  {
	name: string
}

@Component({
    selector: 'app-todo-form',
    templateUrl: './todo-form.component.html',
    styleUrls: ['./todo-form.component.css']
})
export class TodoFormComponent implements OnInit {
    inProgress: boolean = false;
	todo: Todo = {name: ''}
	message: any = '';
	messageClass: any = 'alert-danger';

    constructor(private location: Location, private http:  HttpClient) { }

    ngOnInit() {
    }

    goBack() {
        this.location.back();
    }

	public submitContact(item: Todo){
		this.todo       = item
		let bodyString  = JSON.stringify(this.todo);
		let headers     = new HttpHeaders({ 'Content-Type': 'application/json' });
		this.inProgress = true;

		this.http.post('/api/todos', bodyString, {headers: headers}).subscribe((res) => {
            this.inProgress = false;
            
            if (res['success']) {
                this.goBack()
            } else {
                this.message      = res['message'];
                this.messageClass = 'alert-danger';
            }
        });
	}
}
