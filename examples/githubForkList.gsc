// This script will list forked github repos, from multiple pages.

// assume access rights have already been granted ...

$i = 1;
FOR p FROM 1 TO 4 ; // loop over first pages
    git = PAGE ("https://github.com/xavier268?tab=repositories&type=fork&page=" + GO p) ;
    SELECT 'a[href][itemprop="name codeRepository"]' AS link FROM git  ;
        // print captured repositories
        tt =  TEXT link ;   
        PRINT $i, "  :  ", tt ;
        $i = $i + 1 ;   
        RETURN tt ;
    

