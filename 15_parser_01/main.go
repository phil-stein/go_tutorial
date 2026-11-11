package main

import "fmt"
import "os"
import "strings"
import "strconv"

type data_shader_paths_t struct{
  basic  string
  pbr    string
  screen string
}

type data_t struct{
  font_size     float32 
  frame_rate    int
  window_title  string
  vsync_enabled bool

  shader_paths data_shader_paths_t
}


func main(){

  fmt.Println( "--- reading file from disk ---" )
  // --- reading file from disk ---
  
  var path string = "../assets/custom_00.format"
  f_data, err := os.ReadFile( path )
  if err != nil {
    fmt.Println( "[ERROR] failed reading file at:", path )
    return
  }
  var txt string = string(f_data) 
  fmt.Println( "txt: --------------------------" )
  fmt.Println( txt )
  fmt.Println( "-------------------------------" )

  tokens := parse_string_to_tokens( &txt )

  for idx, t := range tokens{
    fmt.Println( "t[", idx, "]:", t.token_type.to_string(), "t.val:", t.val )
  }

  var data data_t 
  fmt.Println( "-------------------------------" )
  data.Print()

  parse_tokens( &data, tokens )
  
  fmt.Println( "-------------------------------" )
  data.Print()

}

func (d data_t) Print() {
  fmt.Println( "font_size",     d.font_size )
  fmt.Println( "frame_rate",    d.frame_rate )
  fmt.Println( "window_title",  d.window_title )
  fmt.Println( "vsync_enabled", d.vsync_enabled )
  fmt.Println( "shader_paths {" )
  fmt.Println( "  basic",  d.shader_paths.basic )
  fmt.Println( "  pbr",    d.shader_paths.pbr )
  fmt.Println( "  screen", d.shader_paths.screen )
  fmt.Println( "}" )
}


type Tok_Type int
const (
  Tok_Null = iota
  Tok_Square_Bracket_Open
  Tok_Square_Bracket_Closed
  Tok_Curly_Bracket_Open
  Tok_Curly_Bracket_Closed
  Tok_String 
)
func (t Tok_Type) to_string() string {
  switch t {
    case Tok_Null:                  return "Tok_Null" 
    case Tok_Square_Bracket_Open:   return "Tok_Square_Bracket_Open" 
    case Tok_Square_Bracket_Closed: return "Tok_Square_Bracket_Closed"  
    case Tok_Curly_Bracket_Open:    return "Tok_Curly_Bracket_Open"  
    case Tok_Curly_Bracket_Closed:  return "Tok_Curly_Bracket_Closed" 
    case Tok_String:                return "Tok_String" 
    default:                        return "invalid value for Tok_Type"
  }
}
type token_t struct{
  token_type Tok_Type
  val string
}

func parse_string_to_tokens( txt *string ) []token_t {
  
  var ret []token_t


  var str_builder strings.Builder

  var skip_to_end_of_line bool = false

  var in_string bool = false

  for idx, r := range *txt {

    if skip_to_end_of_line {
      if r == '\n' {
        skip_to_end_of_line = false
      } else {
        continue
      }
    }

    if in_string {
      if r == '"' {
        in_string = false
        var str string = str_builder.String()
        ret = append( ret, token_t{ val:str, token_type:Tok_String } )
        str_builder.Reset()
        fmt.Println( "str:", str )
        continue
      } else if r == '\n' {
        fmt.Println( "[ERROR] string not terminalted via '\"' at end of line" )
        return ret
      }
      str_builder.WriteRune( r )
      continue
    }

    if is_space( r ) {
      continue
    }

    if r == '/' && idx +1 < len(*txt) && (*txt)[idx +1] == '/' {
      skip_to_end_of_line = true
      continue
    }

    if r == '"' {
      in_string = true
      continue
    }else if r == '[' {
      ret = append( ret, token_t{ val:"[", token_type:Tok_Square_Bracket_Open } )
    } else if r == ']' {
      ret = append( ret, token_t{ val:"]", token_type:Tok_Square_Bracket_Closed } )
    } else if r == '{' {
      ret = append( ret, token_t{ val:"{", token_type:Tok_Curly_Bracket_Open } )
    } else if r == '}' {
      ret = append( ret, token_t{ val:"}", token_type:Tok_Curly_Bracket_Closed } )
    }else { // if r != ' ' && r != '\t' && r != '\n' 
      str_builder.WriteRune( r )
      if len(*txt) >= idx +1 {
        if !is_alpha_numeric_extra( rune((*txt)[idx +1]) ) || (*txt)[idx +1] == ']' {
          var str string = str_builder.String()
          ret = append( ret, token_t{ val:str, token_type:Tok_String } )
          str_builder.Reset()
          fmt.Println( "str:", str )
        }
      }
    }
  }

  return ret
}

// checks if numeric aka. if r is 0-9
// or if r is alpha aka. a letter 65-90 and 97-122
// extra bc also allows '_' 
func is_alpha_numeric_extra( r rune ) bool {
  // look at https://www.ascii-code.com/
  // runes are technically utf8 but the first 127 runes are the same as ascii
  // 48 - 57 is numeric
  // 65 - 90 is upper case alpha
  // 97 - 122 is lower case alpha
  return ( r >= 48 && r <= 57 ) || ( r >= 65 && r <= 90 ) || ( r >= 97 && r <= 122 ) || ( r == '_' )
}

// checks if space character
// aka. space -> ' ', tab -> '\t', newline -> '\n', etc.
func is_space( r rune ) bool {
  return r <= 32
}


type parser_state_t struct {
  in_shader_paths bool
}

func parse_tokens( data *data_t, tokens []token_t ) {

  var parser_state parser_state_t
  parser_state.in_shader_paths = false

  for idx, t := range tokens{

    if t.token_type == Tok_Curly_Bracket_Closed {
      parser_state.in_shader_paths = false

    // check if [ string ] string, bc. thats a variable assignment
    } else if t.token_type == Tok_Square_Bracket_Open && idx+3 < len(tokens) && 
       tokens[idx +1].token_type == Tok_String && tokens[idx +2].token_type == Tok_Square_Bracket_Closed &&
       tokens[idx +3].token_type == Tok_String {

      parse_tokens_parse_variable_name_value( data, &parser_state, tokens[idx +1].val, tokens[idx +3].val )

    } else if t.token_type == Tok_Square_Bracket_Open && idx+3 < len(tokens) && 
       tokens[idx +1].token_type == Tok_String && tokens[idx +2].token_type == Tok_Square_Bracket_Closed &&
       tokens[idx +3].token_type == Tok_Curly_Bracket_Open {

      parse_tokens_parse_struct_enter( data, &parser_state, tokens[idx +1].val  )
    }
  }

}

func parse_tokens_parse_struct_enter( data *data_t, parser_state *parser_state_t, var_name string  ) {
  if var_name == "shader_paths" {
    (*parser_state).in_shader_paths = true
    return
  }
}

func parse_tokens_parse_variable_name_value( data *data_t, parser_state *parser_state_t, var_name string, value string ) {
  if var_name == "font_size" {
    // fmt.Println( "font_size" )
    v, err := strconv.ParseFloat( value, 64 )
    if err != nil{
      fmt.Println( "[ERROR] parsing value for", var_name, ", value:", value )
      return
    }
    (*data).font_size = float32(v)
  } else if var_name == "frame_rate"{
    // fmt.Println( "frame_rate" )
    v, err := strconv.ParseInt( value, 10, 64 )
    if err != nil{
      fmt.Println( "[ERROR] parsing value for", var_name, ", value:", value )
      return
    }
    (*data).frame_rate = int(v)
  } else if var_name == "window_title"{
    // fmt.Println( "frame_rate" )
    (*data).window_title = strings.Clone(value) 
  } else if var_name == "vsync_enabled"{
    // fmt.Println( "frame_rate" )
    v, err := strconv.ParseBool( value )
    if err != nil{
      fmt.Println( "[ERROR] parsing value for", var_name, ", value:", value )
      return
    }
    (*data).vsync_enabled = bool(v)
  } else if (*parser_state).in_shader_paths &&
            var_name == "basic"{
    // fmt.Println( "frame_rate" )
    (*data).shader_paths.basic = strings.Clone(value) 
  } else if (*parser_state).in_shader_paths &&
            var_name == "pbr"{
    // fmt.Println( "frame_rate" )
    (*data).shader_paths.pbr = strings.Clone(value) 
  } else if (*parser_state).in_shader_paths &&
            var_name == "screen"{
    // fmt.Println( "frame_rate" )
    (*data).shader_paths.screen = strings.Clone(value) 
  } else {
    fmt.Println( "[ERROR] unknown variable name:", var_name )
  }
}

