import { Component, OnInit } from '@angular/core';
import { environment } from '../../environments/environment';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css']
})
export class ProductsComponent implements OnInit {
  
  products = [];
  financing = [];
  confirm = false;
  total = 0;
  successCheckout = false;
  
  fmtPrice = function(value) {
    return 'R$ '+value.toFixed(2).replace('.', ',')
  }
  
  updateTotal = () => {
    let value = 0;
    for (var i = 0; i < this.products.length; i++){
      value = value + (this.products[i].price * this.products[i].quantity)
    }
    this.total = value
  }
  
  onChangeQty = function(e){
    let key = e.target.className
    let prodId = key.split("-")[1];
    this.updateProductQty(prodId, e.target.value)
    this.updateTotal()
  }
  
  updateProductQty = function(prodId, qty){
    for (var i = 0; i < this.products.length; i++){
      // look for the entry with a matching `code` value
      if (this.products[i].id == prodId){
         this.products[i].quantity = qty;
         console.log(this.products[i])
      }
    }
  }
  
  handleConfirm = function(){
    this.confirm = true;
    if(this.total === 0) {
      return
    }
    console.log("A confirm")
    fetch("http://localhost:8000/add_budget").then((response) => {
      console.log(response);
      response.json().then((data) => {
          console.log(data);
        });
    });
  }
  
  alertStatus = function(data){
    window.alert(data)
  }
  
  constructor() {

  };

  ngOnInit(): void {
    if(!environment.production) {
      this.products = [
        { id: 1, description: "Módulo Poli 330W", price: 657.3798, quantity: 0},
        { id: 2, description: "Inversor Mono 3KW 220V", price: 3763.59375, quantity: 0},
        { id: 3, description: "Cabo Solar 6MM2 VM", price: 5.32, quantity: 0 },
        { id: 4, description: "Disjuntor Tripolar 63A", price: 23.4348375, quantity: 0},
        { id: 5, description: "Perfis Alumínio 415MM Telhado 4 Placas", price: 144.6571875, quantity: 0},
        { id: 6, description: "Estrutura Solar Telha Ondulada 4 Placas", price: 165.615625, quantity: 0}
      ]
      
      this.financing = [
        { id: 1, installments_qty: 36, tax: 1.46 },
        { id: 2, installments_qty: 60, tax: 1.45 },
        { id: 3, installments_qty: 72, tax: 1.44 },
        { id: 4, installments_qty: 120, tax: 1.39 },
      ]
    }
  }

}
