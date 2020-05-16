import { Component, OnInit } from '@angular/core';
import { environment } from '../../environments/environment';
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable()
@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css']
})
export class ProductsComponent implements OnInit {
  products = [];
  installments = [];
  confirm = false;
  total = 0;
  tax = 0;
  installmentQty = 0;
  successCheckout = false;
  serverError = false;
  result = {
    installment_value: 0,
    installment_qty: 0,
    tax: 0
  }
  showResult = false;
  
  fmtPrice = function(value) {
    return 'R$ '+value.toFixed(2).replace('.', ',')
  }
  
  fmtPercent = function(value) {
    return value.toFixed(2).replace('.', ',')+'%'
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
      if (this.products[i].id == prodId){
         this.products[i].quantity = qty;
      }
    }
  }
  
  onChangeInstallment = function(e){
    for (var i = 0; i < this.installments.length; i++){
      if (this.installments[i].installments_qty == e.target.value){
         this.tax = this.installments[i].tax;
      }
    }
    this.installmentQty = e.target.value
  }
  
  setTax(){
    if(this.installmentQty === 36){
      this.tax = 1.46
    }
    if(this.installmentQty === 60) {
      this.tax = 1.45
    }
    if(this.installmentQty === 72) {
      this.tax = 1.44
    }
    if(this.installmentQty === 120){
      this.tax = 1.39
    }
  }
  
  handleSubmit = function(){
    this.confirm = true;
    if(this.total === 0) {
      return
    }
    
    this.setTax()
    var obj;
    const rawResponse = fetch('http://localhost:8000/add_budget', {
      method: 'POST',
      body: JSON.stringify({
        "Total": this.total, 
        "Tax": this.tax, 
        "Installment": this.installmentQty
      })
    }).then(res => res.json())
    .then(data => obj = data)
    .then(() => this.setShowTestTrue(obj)).catch(function(err){ window.alert("Falha ao conectar com serviço")}  )
  }
  
  setShowTestTrue(obj) {
    console.log(obj)
    this.result = {
      installment_value: obj.InstallmentValue,
      installment_qty: obj.Installment,
      tax: obj.Tax
    }
    this.showResult = true
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
      
      this.installments = [
        { id: 1, installments_qty: 36, tax: 1.46 },
        { id: 2, installments_qty: 60, tax: 1.45 },
        { id: 3, installments_qty: 72, tax: 1.44 },
        { id: 4, installments_qty: 120, tax: 1.39 }
      ]
    }
  }

}
