#include <AppKit/AppKit.h>
#import <Cocoa/Cocoa.h>
#import <Foundation/Foundation.h>
#include <Foundation/NSObjCRuntime.h>
#include <_string.h>

const char *CheckSpelling(const char *word) {
  NSString *nsWord = [NSString stringWithUTF8String:word];
  NSSpellChecker *spellChecker = [NSSpellChecker sharedSpellChecker];

  NSRange misspelledRange = [spellChecker checkSpellingOfString:nsWord
                                                     startingAt:0];

  if (misspelledRange.location == NSNotFound) {
    return NULL;
  }

  NSArray *guesses = [spellChecker guessesForWordRange:misspelledRange
                                              inString:nsWord
                                              language:nil
                                inSpellDocumentWithTag:0];

  if (guesses.count > 0) {
    return strdup([[guesses firstObject] UTF8String]);
  } else {
    return strdup("No suggestion");
  }
}
