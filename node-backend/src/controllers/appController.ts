import type { Request, Response, NextFunction} from 'express';

import { items, type Item } from '../models/item.js';

export const createItem = (req: Request, res: Response, next: NextFunction) => {
    try {
const { name } = req.body;
        const newItem: Item = { id: Date.now(), name}
        items.push(newItem);
        res.status(201).json(newItem);
    }catch(error) {
        next(error)
    }
};

export const getItems = (req: Request, res: Response, next: NextFunction) => {
  try {
      res.json(items)
  } catch (error) {
    next(error)
  }
};

export const getItemById = (req: Request, res: Response, next: NextFunction) => {
  try {
    const id = parseInt(req.params.id as string, 10);
    const item = items.find((i) => i.id === id);

    if(!item){
      res.status(404).json({ message: 'Item not found'});
      return;
    }
    res.json(item);
  } catch(error){
    next(error)
  }
}

export const updateItem = (req: Request, res: Response, next: NextFunction) => {
  try {
    const id = parseInt(req.params.id as string, 10);
    const { name } = req.body;
    const itemIndex = items.findIndex((i) => i.id === id);
    if(itemIndex === -1){
      res.status(404).json({ message: 'Item not found'});
      return;
    }
  }catch(error){
    next(error)
  }
};

export const deleteItem = (req: Request, res: Response, next: NextFunction) => {
  try {
   const id = parseInt(req.params.id as string, 10);
   const itemIndex = items.findIndex((i) => i.id === id);
   if(itemIndex === -1) {
      res.status(404).json({ message: 'Item not found'});
      return;
   }

   const deletedItem = items.splice(itemIndex, 1)[0];
   res.json(deletedItem)
  }catch(error){
    next(error)
  }
};
