package type5

import (

	     "time"
	     "cloud.google.com/go/storage"
	                                   )


type  Ds_Copy_List    struct           {

          Id             int64
	      Basic_Name     string
	      Copy_Name      string   //   the datastore name which is copied
	      New_Name       string   //  new datastore's name

}


type  Db_Access_List    struct           {

          Id              int64
          Line_No         int64
          Db_Type         string      // datastore's type
                                      //  ds : datastore
                                      //  sr : storage

          Access_Type        string   // access type
                                      //  rename

          Project_Name     string     // project-name
	      Bucket_Name      string     // bucket-name
	      Basic_File_Name    string   // basic file name
	      New_File_Name      string   // new file name
}

///
///    D.B access list
///

type  Db_Access_List2    struct           {

          Id              int64       //　data-id
          Line_No         int64       //
          Db_Type         string      //
                                      //

          Access_Type        string   // access type
                                      //  rename

          Project_Name     string     // プロジェクト名
	      Bucket_Name     string      //
	      Basic_File_Name    string   //
	      New_File_Name      string   //

}

///
///           record for csv file
///

type  Csv_Inf    struct           {

          Id            int64
          Line_No       int64    //
          File_Name     string   //
          Column_Num    int64    // line number
	      Column1       string   // line1
	      Column2       string   //
	      Column3       string   //
	      Column4       string   //
	      Column5       string   //
	      Column6       string   //
	      Column7       string   //
	      Column8       string   //
	      Column9       string   //
	      Column10      string   //
}

///
///
///

type  Interpret    struct           {

//          Id            int64
//          Line_No       int64    // 行NO.

	      Ex_All        string   // 計算式ALL
	      Ex_Parts      string   // 計算式Parts

}
///
///           record for csv file
///

type  Csv_Records    struct           {

      Records_Num    int64   // csv-record-number

//          Id            int64
//          Line_No       int64       // line NO.
      Records[10]    []Csv_Inf   // csv-record
}

///
///   画像ファイル表示
///

type  Image_Show    struct           {

          Id              int64    // data-id
	      File_Name       string   //
	      Url             string   // url
}
///
///     storage-inf.  ( for view )
///

type  Storage_B_O_View    struct           {

          Line_No         int64    // 行NO.
          Project_Name     string   // プロジェクト名
	      Bucket_Name     string   //
	      Object_Name     string
	      Created       time.Time  // the time whitch was created

}

///
///     storage-inf. ( for temp. )
///

type  Storage_B_O_Temp    struct           {

          Id              int64    //
          Line_No         int64    // 行NO.
          Project_Name     string   // プロジェクト名
	      Bucket_Name     string
	      Object_Name     string   // オブジェクト名
	      Created       time.Time  //  create-time

}

///
///     storage-inf.
///

type  Storage_B_O    struct           {

          Id              int64
          Line_No         int64    // 行NO.
          Project_Name     string   // プロジェクト名
	      Bucket_Name     string   //
	      Object_Name     string
	      Created       time.Time  //   create-time

}
///
///    for work
///

type  General_Work    struct           {

          Int64_Work     int64           // int型ワークエリア
          Float64_Work   float64         // float型ワークエリア
	      String_Work    string          // string型ワークエリア
	      Sw_Work        *storage.Writer // storage-writer
}
