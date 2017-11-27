import { Component, OnInit } from '@angular/core';
import { HttpClient, HttpHeaders, HttpRequest } from '@angular/common/http';

@Component({
  selector: 'app-todo-list',
  templateUrl: './todo-list.component.html',
  styleUrls: ['./todo-list.component.css']
})
export class TodoListComponent implements OnInit {
    items:any ;
    itemsModels:any = [];

    constructor(private http: HttpClient) { }

    ngOnInit() {
        this.http.get("/api/todos").subscribe(res => {
            this.items = res;
        });
    }

    markCompleted(index: any){
		let bodyString  = JSON.stringify({"completed": this.items[index].completed});
		let headers     = new HttpHeaders({ 'Content-Type': 'application/json' });

		this.http.put(`/api/todos/${this.items[index].id}`, bodyString, {headers: headers}).subscribe((res) => {
            if (!res['success']) 
                alert(res['message']); 
        });
    }

    deleteItem(id: any) {
        this.http.delete(`/api/todos/${id}`).subscribe(res => {
            if (res['success']) {
                this.items = res['todos'];
            } else {
                alert(res['message']);
            }
        });
    }
}
