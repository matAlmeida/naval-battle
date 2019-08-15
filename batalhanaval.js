tiro(x,y)
{
if (mat[x][y] == "agua")
    { lista.remove(mat(x,y)) }

    if (mat[x][y] == "submarino")
        {
            lista.remove(mat(x,y))
            lista.remove(mat(x+1,y)) //cruz
            lista.remove(mat(x-1,y)) //cruz
            lista.remove(mat(x,y-1)) //cruz
            lista.remove(mat(x,y+1)) //cruz
            lista.remove(mat(x-1,y-1)) //Diagonais 
            lista.remove(mat(x+1,y+1))  //Diagonais 
            lista.remove(mat(x-1,y+1)) //Diagonais 
            lista.remove(mat(x+1,y-1)) //Diagonais 

            submarino.remove();
        }

        if (mat[x][y] == "destroier")
            {

                lista.remove(mat(x,y))
                lista.remove(mat(x-1,y-1)) //remove diagonais 
                lista.remove(mat(x+1,y+1)) //remove diagonais
                lista.remove(mat(x-1,y+1)) //remove diagonais
                lista.remove(mat(x+1,y-1)) //remove diagonais

                lista.atira(mat(x+1,y)) //sugestão de tiro 
                lista.atira(mat(x-1,y)) //sugestão de tiro
                lista.atira(mat(x,y-1)) //sugestão de tiro
                lista.atira(mat(x,y+1)) //sugestão de tiro


            }

         
            if (mat[x][y] == "cruzadores")
            {

                
                lista.remove(mat(x,y))
                lista.remove(mat(x-1,y-1)) //remove diagonais 
                lista.remove(mat(x+1,y+1)) //remove diagonais
                lista.remove(mat(x-1,y+1)) //remove diagonais
                lista.remove(mat(x+1,y-1)) //remove diagonais

                lista.atira(mat(x+1,y)) //sugestão de tiro 
                lista.atira(mat(x-1,y)) //sugestão de tiro
                lista.atira(mat(x,y-1)) //sugestão de tiro
                lista.atira(mat(x,y+1)) //sugestão de tiro
            }

            if (mat[x][y] == "porta-avião")
            {
                lista.remove(mat(x,y))
                lista.remove(mat(x-1,y-1)) //remove diagonais 
                lista.remove(mat(x+1,y+1)) //remove diagonais
                lista.remove(mat(x-1,y+1)) //remove diagonais
                lista.remove(mat(x+1,y-1)) //remove diagonais

                lista.atira(mat(x+1,y)) //sugestão de tiro 
                lista.atira(mat(x-1,y)) //sugestão de tiro
                lista.atira(mat(x,y-1)) //sugestão de tiro
                lista.atira(mat(x,y+1)) //sugestão de tiro 
            }

            if (mat[x][y] == "hidro-avião")
            {
                lista.remove(mat(x+1,y)) // remove cruz
                lista.remove(mat(x-1,y)) //remove cruz
                lista.remove(mat(x,y-1)) //remove cruz
                lista.remove(mat(x,y+1)) //remove cruz


                lista.remove(mat(x-1,y-1))  //sugestão de tiro 
                lista.remove(mat(x+1,y+1))  //sugestão de tiro 
                lista.remove(mat(x-1,y+1))  //sugestão de tiro 
                lista.remove(mat(x+1,y-1))  //sugestão de tiro 


            }
    

    
}



