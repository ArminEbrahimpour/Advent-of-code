use std::{fs};

fn Read_file()-> String{

    let contents = fs::read_to_string("./input.txt").expect("shoud read the file ");

    contents

}


fn main() {

    let data = Read_file();

    println!("{}",&data);

}
