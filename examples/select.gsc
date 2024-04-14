// Select specific elements in a page

page = PAGE "http://www.wikipedia.fr" ;
$i = 0 ;
SELECT "a" AS loop FROM page LIMIT 5;
    $i = $i + 1 ;
    RETURN $i, TEXT loop ;
